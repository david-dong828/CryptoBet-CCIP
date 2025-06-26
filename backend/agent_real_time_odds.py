#!/usr/bin/env python3

import os
import json
import time
import re
import logging

from contextlib import contextmanager, suppress
from dataclasses import dataclass, asdict, field
from datetime import datetime
from typing import List, Dict, Optional, Set, Iterator, Union
from pathlib import Path
import sys

from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.common.by import By
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.common.exceptions import WebDriverException, TimeoutException

import openai
import schedule
from dotenv import load_dotenv

load_dotenv('bk_api.env')

# Configure clean logging and suppress ALL unwanted output
logging.basicConfig(level=logging.INFO, format='%(message)s')
logger = logging.getLogger(__name__)

# Suppress Selenium warnings completely
logging.getLogger('selenium').setLevel(logging.CRITICAL)
logging.getLogger('urllib3').setLevel(logging.CRITICAL)
logging.getLogger('selenium.webdriver.remote.remote_connection').setLevel(logging.CRITICAL)

# Redirect stderr to suppress the service termination error
class ErrorSuppressor:
    def __init__(self):
        self.original_stderr = sys.stderr
        
    def __enter__(self):
        sys.stderr = open(os.devnull, 'w')
        return self
        
    def __exit__(self, exc_type, exc_val, exc_tb):
        sys.stderr.close()
        sys.stderr = self.original_stderr


class Config:
    """Centralized configuration"""
    BASE_URL = "https://polymarket.com/sports/soccer/props"
    OUTPUT_DIR = Path("outputs")
    
    # Selenium settings (from your working code)
    SCROLL_ATTEMPTS = 5
    SCROLL_DELAY = 3
    PAGE_LOAD_DELAY = 5
    LINK_COLLECTION_ATTEMPTS = 3
    MIN_PREDICTIONS = 2
    
    # Chrome paths and options (from your working code)
    CHROME_PATHS = [
        "/usr/bin/chromium-browser", 
        "/usr/bin/google-chrome",
        "/usr/bin/chrome"
    ]
    
    CHROME_OPTIONS = [
        "--headless=new",
        "--no-sandbox",
        "--disable-dev-shm-usage",
        "--disable-logging",
        "--log-level=3",
        "--disable-gpu",
        "--disable-extensions",
        "--disable-default-apps",
        "--disable-background-timer-throttling",
        "--disable-backgrounding-occluded-windows",
        "--disable-renderer-backgrounding",
        "--disable-features=TranslateUI,VizDisplayCompositor",
        "--no-first-run",
        "--no-default-browser-check",
        "--disable-ipc-flooding-protection",
        "--remote-debugging-port=0",
        "--silent",
        "--disable-crash-reporter",
    ]
    
    # OpenAI settings
    OPENAI_API_KEY = os.getenv("OPENAI_API_KEY")
    OPENAI_MODEL = "gpt-4o-mini"
    
    @classmethod
    def validate(cls) -> None:
        """Validate configuration"""
        if not cls.OPENAI_API_KEY:
            logger.warning("⚠️ OPENAI_API_KEY not set - using fallback extraction only")


@dataclass(frozen=True)
class Prediction:
    """Immutable prediction data (from your original code)"""
    team_name: str
    probability: float
    
    def __post_init__(self):
        if not (0 <= self.probability <= 1):
            raise ValueError(f"Probability must be between 0 and 1, got {self.probability}")
    
    @property
    def dict_key(self) -> str:
        """Convert team name to dictionary key format (from your original)"""
        key = self.team_name.lower().replace(' ', '_').replace('&', 'and')
        key = re.sub(r'[^\w]', '_', key)
        return f"{key}_win"


@dataclass
class Competition:
    """Competition data container (from your original code)"""
    match: str
    date: str
    polymarket_prediction: Dict[str, float] = field(default_factory=dict)
    
    @classmethod
    def from_predictions(cls, match: str, predictions: List[Prediction]) -> 'Competition':
        """Create competition from list of predictions"""
        pred_dict = {pred.dict_key: pred.probability for pred in predictions}
        
        shortened_match = match
        if ':' in match:
            shortened_match = match.split(':', 1)[1].strip()
            shortened_match = shortened_match.replace("vs.", "vs").replace("VS.", "vs").strip()

            return cls(
                match=shortened_match,
                date=datetime.now().strftime("%Y-%m-%d"),
                polymarket_prediction=pred_dict
            )

    
    def to_dict(self) -> Dict:
        """Convert to dictionary for serialization"""
        return asdict(self)


class CompetitionFilter:
    """Smart filtering for competitions (from your original logic)"""
    
    SKIP_TERMS = frozenset([
        'ballon d\'or', 'golden ball', 'top goalscorer', 'most goals',
        'goals scored', 'club world cup', 'fifa club world cup',
        'will score', 'to score', 'transfer', 'sign with'
    ])
    
    INCLUDE_TERMS = frozenset([
        '2026 fifa world cup', '2026 world cup', 'world cup',
        'gold cup', 'concacaf', 'uefa', 'european u21', 'u21', 'u19',
        'qualify', 'qualification', 'vs.', 'vs ', ' vs', ' v ', 'against',
        'copa america', 'nations league', 'championship'
    ])
    
    @classmethod
    def should_process(cls, title: str) -> bool:
        """Determine if competition should be processed"""
        title_lower = title.lower()
        
        if any(term in title_lower for term in cls.SKIP_TERMS):
            return False
        
        return any(term in title_lower for term in cls.INCLUDE_TERMS)


class TextProcessor:
    """Advanced text processing utilities (from your original code)"""
    
    INVALID_TERMS = frozenset({
        'yes', 'no', 'buy', 'sell', 'outcome', 'chance', 'vol', 'volume', 'new'
    })
    
    CLEANING_PATTERNS = [
        (re.compile(r'\d+(?:\.\d+)?\s*%.*'), ''),
        (re.compile(r'Buy\s+(Yes|No).*'), ''),
        (re.compile(r'\$.*Vol\.'), ''),
        (re.compile(r'\d+(?:\.\d+)?¢'), ''),
        (re.compile(r'\bNEW\b'), ''),
        (re.compile(r'\d{4}-\d{2}-\d{2}'), ''),
        (re.compile(r'[^\w\s]'), ' '),
    ]
    
    @classmethod
    def clean_text(cls, text: str) -> str:
        """Clean text using compiled regex patterns"""
        if not text:
            return ""
        
        clean_text = text.strip()
        for pattern, replacement in cls.CLEANING_PATTERNS:
            clean_text = pattern.sub(replacement, clean_text).strip()
        
        clean_text = re.sub(r'\s+', ' ', clean_text)
        return clean_text
    
    @classmethod
    def is_valid_team_name(cls, text: str) -> bool:
        """Validate team name with stricter criteria"""
        if not text or len(text) < 2 or len(text) > 50:
            return False
        
        text_lower = text.lower()
        if any(term in text_lower for term in cls.INVALID_TERMS):
            return False
        
        if not re.search(r'[a-zA-Z]', text):
            return False
        
        return True


class SafeDriverManager:
    """Enhanced driver manager (from your working code)"""
    
    def __init__(self, headless: bool = True):
        self.headless = headless
        self.driver = None
        self.service = None
    
    def __enter__(self) -> webdriver.Chrome:
        """Create and return Chrome driver"""
        self.driver = self._create_driver()
        return self.driver
    
    def __exit__(self, exc_type, exc_val, exc_tb):
        """Clean exit without throwing exceptions"""
        self._safe_cleanup()
    
    def _create_driver(self) -> webdriver.Chrome:
        """Create Chrome driver with optimal settings"""
        options = self._get_chrome_options()
        
        # Try without service first to avoid termination issues
        try:
            return webdriver.Chrome(options=options)
        except Exception:
            # Fallback to service but override cleanup
            try:
                self.service = Service(log_output=os.devnull)
                # Override the service terminate method to prevent errors
                original_terminate = self.service.stop
                self.service.stop = lambda: None
                return webdriver.Chrome(service=self.service, options=options)
            except Exception:
                raise RuntimeError("Failed to create Chrome driver")
    
    def _get_chrome_options(self) -> Options:
        """Configure Chrome options for stability"""
        options = Options()
        
        for option in Config.CHROME_OPTIONS:
            options.add_argument(option)
        
        for path in Config.CHROME_PATHS:
            if Path(path).exists():
                options.binary_location = path
                break
        
        options.add_experimental_option('excludeSwitches', ['enable-logging'])
        options.add_experimental_option('useAutomationExtension', False)
        
        return options
    
    def _safe_cleanup(self):
        """Cleanup without throwing exceptions"""
        with suppress(Exception):
            if self.driver:
                # Disable service cleanup to prevent permission errors
                if hasattr(self.driver, 'service') and self.driver.service:
                    self.driver.service.stop = lambda: None
                    
                with suppress(Exception):
                    self.driver.quit()
        
        # Completely skip service cleanup
        self.service = None


class AIEnhancedExtractor:
    """OpenAI-enhanced prediction extractor"""
    
    def __init__(self):
        self.openai_client = None
        if Config.OPENAI_API_KEY:
            self.openai_client = openai.OpenAI(api_key=Config.OPENAI_API_KEY)
    
    def extract_predictions(self, driver: webdriver.Chrome) -> List[Prediction]:
        """Extract predictions using multiple strategies"""
        predictions = []
        
        # Strategy 1: Percentage element extraction (your most reliable method)
        percentage_predictions = self._extract_percentages(driver)
        predictions.extend(percentage_predictions)
        
        # Strategy 2: Regex extraction (your proven method)
        regex_predictions = self._extract_with_regex(driver)
        predictions.extend(regex_predictions)
        
        # Strategy 3: AI-enhanced extraction (if API available and we need more)
        if self.openai_client and len(predictions) < Config.MIN_PREDICTIONS:
            ai_predictions = self._extract_with_ai(driver)
            predictions.extend(ai_predictions)
        
        # Strategy 4: Extended percentage search for stubborn cases
        if len(predictions) < Config.MIN_PREDICTIONS:
            extended_predictions = self._extract_extended_percentages(driver)
            predictions.extend(extended_predictions)
        
        return list(self._deduplicate_predictions(predictions))
    
    def _extract_with_regex(self, driver: webdriver.Chrome) -> List[Prediction]:
        """Extract using regex patterns (from your RegexExtractor)"""
        countries = [
            'Canada', 'Mexico', 'United States', 'USA', 'Panama', 'Costa Rica',
            'Honduras', 'Saudi Arabia', 'Jamaica', 'El Salvador', 'Guatemala',
            'Haiti', 'Trinidad', 'Tobago', 'Colombia', 'Netherlands', 'Uruguay',
            'Belgium', 'Croatia', 'Austria', 'Sweden', 'Italy', 'Poland',
            'Finland', 'Ukraine', 'Oman', 'England', 'Germany', 'France',
            'Denmark', 'Spain', 'Portugal', 'Brazil', 'Argentina', 'Chile',
            'Peru', 'Ecuador', 'Bolivia', 'Paraguay', 'Venezuela', 'Japan',
            'Korea', 'Australia', 'Morocco', 'Ghana', 'Nigeria', 'Egypt',
            'Norway', 'Switzerland', 'Greece', 'Turkey', 'Serbia', 'Czech',
            'Slovakia', 'Hungary', 'Romania', 'Bulgaria', 'Slovenia', 'Israel',
            'Iran', 'Iraq', 'Jordan', 'Lebanon', 'Qatar', 'Bahrain', 'Kuwait', 'Draw'
        ]
        
        predictions = []
        
        with suppress(Exception):
            page_source = driver.page_source
            countries_pattern = '|'.join(countries)
            
            patterns = [
                re.compile(rf'({countries_pattern})[^%]*?(\d+(?:\.\d+)?)\s*%', re.IGNORECASE),
                re.compile(rf'(\d+(?:\.\d+)?)\s*%[^a-zA-Z]*?({countries_pattern})', re.IGNORECASE)
            ]
            
            for pattern in patterns:
                matches = pattern.findall(page_source)
                for match in matches:
                    with suppress(Exception):
                        country, percent_str = match
                        if not percent_str.replace('.', '').isdigit():
                            country, percent_str = percent_str, country
                        
                        if percent_str.replace('.', '').isdigit():
                            probability = float(percent_str) / 100
                            if 0.001 <= probability <= 1.0 and country in countries:
                                predictions.append(Prediction(country, probability))
        
        return predictions
    
    def _extract_with_ai(self, driver: webdriver.Chrome) -> List[Prediction]:
        """Extract using OpenAI for complex cases"""
        try:
            # Get ALL text elements, not just percentage ones
            all_elements = []
            
            # Try multiple element types
            element_selectors = [
                "//*[contains(text(), '%')]",
                "//div[text()]",
                "//span[text()]", 
                "//p[text()]",
                "//*[@class]"
            ]
            
            for selector in element_selectors:
                with suppress(Exception):
                    elements = driver.find_elements(By.XPATH, selector)
                    all_elements.extend(elements[:10])  # Limit per selector
            
            # Get page source as backup
            page_text = ""
            with suppress(Exception):
                page_text = driver.find_element(By.TAG_NAME, "body").text
            
            text_content = []
            
            # Process elements
            for element in all_elements[:30]:  # Limit total elements
                with suppress(Exception):
                    text = element.text.strip()
                    if text and (('%' in text) or ('¢' in text) or any(c.isdigit() for c in text)):
                        text_content.append(text)
            
            # Add relevant page text chunks
            if page_text:
                chunks = page_text.split('\n')
                for chunk in chunks[:50]:  # Limit chunks
                    if '%' in chunk or any(country in chunk.lower() for country in ['usa', 'mexico', 'canada', 'brazil', 'germany', 'france', 'england', 'italy', 'spain']):
                        text_content.append(chunk.strip())
            
            if not text_content:
                return []
            
            combined_text = '\n'.join(text_content[:100])  # Limit total content
            
            prompt = f"""
            Extract ALL betting predictions from this Polymarket competition data.
            Look for team/country names with percentages, odds, or probability indicators.
            
            IMPORTANT Rules:
            - Extract EVERY team/country with a percentage (like "Brazil 65%" or "45% Germany")
            - Convert percentages to decimals (65% = 0.65, 45% = 0.45)
            - Look for patterns like "TeamName XX%" or "XX% TeamName"
            - Include Draw as a valid option
            - Ignore: yes, no, buy, sell, outcome, chance, vol, volume, new
            - Be aggressive - extract all possible teams with probabilities
            
            Return ONLY JSON array: [{{"team_name": "Brazil", "probability": 0.65}}, {{"team_name": "Germany", "probability": 0.45}}]
            
            Data:
            {combined_text}
            """
            
            response = self.openai_client.chat.completions.create(
                model=Config.OPENAI_MODEL,
                messages=[{"role": "user", "content": prompt}],
                temperature=0,
                max_tokens=1000
            )
            
            result = response.choices[0].message.content.strip()
            
            # Clean response more aggressively
            if '```json' in result:
                start = result.find('```json') + 7
                end = result.find('```', start)
                if end > start:
                    result = result[start:end]
            elif '```' in result:
                start = result.find('```') + 3
                end = result.find('```', start)
                if end > start:
                    result = result[start:end]
            
            # Try to find JSON array even if wrapped
            if '[' in result and ']' in result:
                start = result.find('[')
                end = result.rfind(']') + 1
                result = result[start:end]
            
            pred_data = json.loads(result)
            predictions = []
            
            for item in pred_data:
                try:
                    team_name = str(item['team_name']).strip()
                    probability = float(item['probability'])
                    
                    # More lenient validation for AI results
                    if len(team_name) >= 2 and 0.001 <= probability <= 1.0:
                        # Still check for obvious invalid terms
                        if not any(invalid in team_name.lower() for invalid in ['yes', 'no', 'buy', 'sell', 'vol']):
                            predictions.append(Prediction(team_name, probability))
                
                except (KeyError, ValueError, TypeError):
                    continue
            
            logger.debug(f"AI extracted {len(predictions)} predictions")
            return predictions
        
        except Exception as e:
            logger.debug(f"AI extraction failed: {e}")
            return []
    
    def _extract_percentages(self, driver: webdriver.Chrome) -> List[Prediction]:
        """Extract from percentage elements (from your PercentageExtractor)"""
        predictions = []
        
        with suppress(Exception):
            elements = driver.find_elements(By.XPATH, "//*[contains(text(), '%')]")
            
            for element in elements:
                with suppress(Exception):
                    element_text = element.text.strip()
                    if '%' not in element_text:
                        continue
                    
                    try:
                        parent = element.find_element(By.XPATH, "..")
                        context_text = parent.text.strip()
                    except:
                        context_text = element_text
                    
                    percent_matches = re.findall(r'(\d+(?:\.\d+)?)\s*%', context_text)
                    
                    for percent_str in percent_matches:
                        try:
                            probability = float(percent_str) / 100
                            if not (0.001 <= probability <= 1.0):
                                continue
                            
                            team_name = self._extract_team_name(context_text, percent_str)
                            if team_name:
                                predictions.append(Prediction(team_name, probability))
                        
                        except (ValueError, TypeError):
                            continue
        
        return predictions
    
    def _extract_team_name(self, context_text: str, percent_str: str) -> Optional[str]:
        """Extract team name from context (from your original)"""
        lines = context_text.split('\n')
        
        percent_line_idx = next(
            (i for i, line in enumerate(lines) if f"{percent_str}%" in line),
            -1
        )
        
        if percent_line_idx >= 0:
            start = max(0, percent_line_idx - 2)
            end = min(len(lines), percent_line_idx + 3)
            search_lines = lines[start:end]
        else:
            search_lines = lines
        
        for line in search_lines:
            clean_line = TextProcessor.clean_text(line)
            if TextProcessor.is_valid_team_name(clean_line):
                return clean_line
        
    def _extract_extended_percentages(self, driver: webdriver.Chrome) -> List[Prediction]:
        """Extended percentage extraction for stubborn cases"""
        predictions = []
        
        with suppress(Exception):
            # Try different selectors to find percentage elements
            selectors = [
                "//*[contains(text(), '%')]",
                "//div[contains(text(), '%')]",
                "//span[contains(text(), '%')]", 
                "//p[contains(text(), '%')]",
                "//*[contains(@class, 'percentage')]",
                "//*[contains(@class, 'prob')]",
                "//*[contains(@class, 'odds')]"
            ]
            
            all_elements = set()
            for selector in selectors:
                with suppress(Exception):
                    elements = driver.find_elements(By.XPATH, selector)
                    all_elements.update(elements)
            
            for element in all_elements:
                with suppress(Exception):
                    # Get broader context
                    contexts = []
                    
                    # Try element text
                    contexts.append(element.text.strip())
                    
                    # Try parent text
                    with suppress(Exception):
                        parent = element.find_element(By.XPATH, "..")
                        contexts.append(parent.text.strip())
                    
                    # Try grandparent text
                    with suppress(Exception):
                        grandparent = element.find_element(By.XPATH, "../..")
                        contexts.append(grandparent.text.strip())
                    
                    for context_text in contexts:
                        if not context_text or '%' not in context_text:
                            continue
                        
                        # Extract all percentages from context
                        percent_matches = re.findall(r'(\d+(?:\.\d+)?)\s*%', context_text)
                        
                        for percent_str in percent_matches:
                            try:
                                probability = float(percent_str) / 100
                                if not (0.001 <= probability <= 1.0):
                                    continue
                                
                                # Try multiple team name extraction methods
                                team_names = self._extract_all_team_names(context_text, percent_str)
                                
                                for team_name in team_names:
                                    if team_name:
                                        predictions.append(Prediction(team_name, probability))
                                        break  # Only take first valid team name
                            
                            except (ValueError, TypeError):
                                continue
        
        return predictions
    
    def _extract_all_team_names(self, context_text: str, percent_str: str) -> List[str]:
        """Try multiple methods to extract team names"""
        team_names = []
        
        lines = context_text.split('\n')
        
        # Method 1: Look around percentage line (your original method)
        percent_line_idx = next(
            (i for i, line in enumerate(lines) if f"{percent_str}%" in line),
            -1
        )
        
        if percent_line_idx >= 0:
            start = max(0, percent_line_idx - 2)
            end = min(len(lines), percent_line_idx + 3)
            search_lines = lines[start:end]
        else:
            search_lines = lines
        
        for line in search_lines:
            clean_line = TextProcessor.clean_text(line)
            if TextProcessor.is_valid_team_name(clean_line):
                team_names.append(clean_line)
        
        # Method 2: Look for country/team names in full context
        country_patterns = [
            r'\b(Canada|Mexico|United States|USA|Panama|Costa Rica|Honduras|Saudi Arabia|Jamaica|El Salvador|Guatemala|Haiti|Trinidad|Tobago|Colombia|Netherlands|Uruguay|Belgium|Croatia|Austria|Sweden|Italy|Poland|Finland|Ukraine|Oman|England|Germany|France|Denmark|Spain|Portugal|Brazil|Argentina|Chile|Peru|Ecuador|Bolivia|Paraguay|Venezuela|Japan|Korea|Australia|Morocco|Ghana|Nigeria|Egypt|Norway|Switzerland|Greece|Turkey|Serbia|Czech|Slovakia|Hungary|Romania|Bulgaria|Slovenia|Israel|Iran|Iraq|Jordan|Lebanon|Qatar|Bahrain|Kuwait|Draw)\b',
            r'\b([A-Z][a-z]+ [A-Z][a-z]+)\b',  # Two-word team names
            r'\b([A-Z][a-z]{2,})\b'  # Single country names
        ]
        
        for pattern in country_patterns:
            matches = re.findall(pattern, context_text, re.IGNORECASE)
            for match in matches:
                team_name = match if isinstance(match, str) else match[0] if match else ""
                if TextProcessor.is_valid_team_name(team_name):
                    team_names.append(team_name)
        
        # Method 3: Parse structured text (lines before/after percentage)
        for i, line in enumerate(lines):
            if f"{percent_str}%" in line:
                # Check previous lines
                for j in range(max(0, i-3), i):
                    candidate = TextProcessor.clean_text(lines[j])
                    if TextProcessor.is_valid_team_name(candidate):
                        team_names.append(candidate)
                
                # Check next lines  
                for j in range(i+1, min(len(lines), i+4)):
                    candidate = TextProcessor.clean_text(lines[j])
                    if TextProcessor.is_valid_team_name(candidate):
                        team_names.append(candidate)
        
        return team_names
    
    def _deduplicate_predictions(self, predictions: List[Prediction]) -> Iterator[Prediction]:
        """Remove duplicate predictions (from your original logic)"""
        seen_teams = {}
        
        for pred in predictions:
            key = pred.dict_key
            if key not in seen_teams or pred.probability > seen_teams[key].probability:
                seen_teams[key] = pred
        
        yield from seen_teams.values()


class PolymarketScraper:
    """Main scraper orchestrator using your proven Selenium approach"""
    
    def __init__(self, headless: bool = True):
        self.headless = headless
        self.extractor = AIEnhancedExtractor()
    
    def scrape_all_competitions(self) -> List[Competition]:
        """Main entry point for scraping (using your working logic)"""
        logger.info("🚀 Starting Polymarket Competition Scraper")
        
        try:
            # Use error suppressor to hide service termination errors
            with ErrorSuppressor():
                with SafeDriverManager(self.headless) as driver:
                    return self._execute_scraping(driver)
        except Exception as e:
            logger.error(f"❌ Scraping failed: {e}")
            return []
    
    def _execute_scraping(self, driver: webdriver.Chrome) -> List[Competition]:
        """Execute the scraping process (from your working code)"""
        competitions = []
        
        # Get event links using your proven method
        event_links = self._collect_event_links(driver)
        filtered_links = [
            (url, title) for url, title in event_links 
            if CompetitionFilter.should_process(title)
        ]
        
        logger.info(f"📊 Processing {len(filtered_links)} relevant competitions")
        
        # Process each competition
        for i, (url, title) in enumerate(filtered_links, 1):
            logger.info(f"🏆 [{i}/{len(filtered_links)}] {title}")
            
            with suppress(Exception):
                competition = self._scrape_single_competition(driver, url, title)
                if competition:
                    competitions.append(competition)
                    pred_count = len(competition.polymarket_prediction)
                    logger.info(f"   ✅ Extracted {pred_count} predictions")
                else:
                    logger.info("   ⚠️ No predictions found")
        
        return competitions
    
    def _collect_event_links(self, driver: webdriver.Chrome) -> List[tuple]:
        """Collect all event links (your proven working method)"""
        logger.info("🌐 Loading main page...")
        driver.get(Config.BASE_URL)
        time.sleep(Config.PAGE_LOAD_DELAY)
        
        # Scroll to load content (your method)
        logger.info("📜 Loading all content...")
        for i in range(Config.SCROLL_ATTEMPTS):
            driver.execute_script("window.scrollTo(0, document.body.scrollHeight);")
            time.sleep(Config.SCROLL_DELAY)
        
        # Collect links using your exact method
        all_links = set()
        
        for attempt in range(Config.LINK_COLLECTION_ATTEMPTS):
            with suppress(Exception):
                links = driver.find_elements(By.CSS_SELECTOR, "a[href*='/event/']")
                
                for link in links:
                    with suppress(Exception):
                        url = link.get_attribute("href")
                        title = link.text.strip()
                        
                        if url and title and len(title) > 3:
                            all_links.add((url, title))
        
        logger.info(f"📋 Found {len(all_links)} unique events")
        return list(all_links)
    
    def _scrape_single_competition(self, driver: webdriver.Chrome, url: str, title: str) -> Optional[Competition]:
        """Scrape individual competition (enhanced with AI)"""
        driver.get(url)
        time.sleep(Config.PAGE_LOAD_DELAY)
        
        # Get page title
        page_title = self._extract_page_title(driver, title)
        
        # Extract predictions using AI-enhanced extractor
        predictions = self.extractor.extract_predictions(driver)
        
        if len(predictions) >= Config.MIN_PREDICTIONS:
            return Competition.from_predictions(page_title, predictions)
        
        return None
    
    def _extract_page_title(self, driver: webdriver.Chrome, fallback: str) -> str:
        """Extract page title with fallback (from your original)"""
        with suppress(Exception):
            element = driver.find_element(By.TAG_NAME, "h1")
            title = element.text.strip()
            if title:
                return title
        
        return fallback


class ResultManager:
    """Handle result storage and management (from your original code)"""
    
    @staticmethod
    def save_competitions(competitions: List[Competition]) -> Optional[Path]:
        """Save competitions to JSON files"""
        if not competitions:
            logger.error("❌ No competitions to save")
            return None
        
        Config.OUTPUT_DIR.mkdir(exist_ok=True)
        
        data = [comp.to_dict() for comp in competitions]
        
        filename = Config.OUTPUT_DIR / f"country_competitions.json"
        
        ResultManager._write_json(filename, data)
        
        latest_file = Config.OUTPUT_DIR / "latest_country_competitions.json"
        ResultManager._write_json(latest_file, data)
        
        logger.info(f"💾 Saved {len(competitions)} competitions to {filename}")
        return filename
    
    @staticmethod
    def _write_json(filepath: Path, data: List[Dict]):
        """Write data to JSON file"""
        with open(filepath, "w", encoding="utf-8") as f:
            json.dump(data, f, indent=2, ensure_ascii=False)


def main():
    """Main application entry point"""
    try:
        Config.validate()
        scraper = PolymarketScraper(headless=True)
        competitions = scraper.scrape_all_competitions()
        
        if competitions:
            filename = ResultManager.save_competitions(competitions)
            logger.info(f"\n✅ Successfully scraped {len(competitions)} competitions")
            logger.info(f"📁 Data saved to: {filename}")
        else:
            logger.info("\n❌ No competitions found")
    
    except KeyboardInterrupt:
        logger.info("\n⚠️ Process interrupted by user")
    except Exception as e:
        logger.error(f"\n❌ Unexpected error: {e}")
    finally:
        logger.info("🔒 Process complete")

def job():
    main()

if __name__ == "__main__":
    # Schedule the job to run every 30 minutes
    schedule.every(15).minutes.do(job)

    logger.info("⏱️ Real-time odds Agent will run every 15 minutes.")
    try:
        while True:
            schedule.run_pending()
            time.sleep(1)
    except KeyboardInterrupt:
        logger.info("\n🛑 Scheduler interrupted by user")