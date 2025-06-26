import pandas as pd
from sklearn.model_selection import train_test_split
from sklearn.ensemble import RandomForestClassifier
import joblib
import numpy as np
import os

FOLDER = os.path.dirname(__file__)
TRAINING_FILE = os.path.join(FOLDER,'football_training_data/results.csv') 

# ----------------------------
# 1. Data Preparation Module
# ----------------------------
def load_and_prepare_data(csv_path):
    df = pd.read_csv(csv_path)
    df['date'] = pd.to_datetime(df['date'])
    df["home_win"] = (df["home_score"] > df["away_score"]).astype(int)
    return df

# ----------------------------
# 2. Feature Engineering Module
# ----------------------------
def compute_rolling_features(df, team_col, prefix):
    win_rate, goals_for, goals_against, last_game_days = [], [], [], []
    last_games = {}

    for idx, row in df.iterrows():
        team = row[team_col]
        date = row['date']

        prev_games = df[((df['home_team'] == team) | (df['away_team'] == team)) & (df['date'] < date)].sort_values('date').tail(10)
        wins = 0
        gf = 0
        ga = 0
        for _, g in prev_games.iterrows():
            if g['home_team'] == team:
                gf += g['home_score']
                ga += g['away_score']
                if g['home_score'] > g['away_score']:
                    wins += 1
            else:
                gf += g['away_score']
                ga += g['home_score']
                if g['away_score'] > g['home_score']:
                    wins += 1
        games_played = len(prev_games)
        win_rate.append(wins / games_played if games_played > 0 else 0.5)
        goals_for.append(gf / games_played if games_played > 0 else 1.0)
        goals_against.append(ga / games_played if games_played > 0 else 1.0)

        if team in last_games:
            days_since = (date - last_games[team]).days
        else:
            days_since = 7
        last_game_days.append(days_since)
        last_games[team] = date

    df[f'{prefix}_win_rate'] = win_rate
    df[f'{prefix}_goals_for'] = goals_for
    df[f'{prefix}_goals_against'] = goals_against
    df[f'{prefix}_days_since'] = last_game_days
    return df

def add_features(df):
    df = compute_rolling_features(df, 'home_team', 'home')
    df = compute_rolling_features(df, 'away_team', 'away')
    df = pd.concat([df, pd.get_dummies(df['tournament'], prefix='tournament')], axis=1)
    df["neutral"] = df["neutral"].astype(int)
    return df

# ----------------------------
# 3. Model Training Module
# ----------------------------
def train_model(df):
    feature_cols = [
        "home_win_rate", "away_win_rate",
        "home_goals_for", "home_goals_against", "away_goals_for", "away_goals_against",
        "home_days_since", "away_days_since",
        "neutral"
    ] + [col for col in df.columns if col.startswith("tournament_")]
    X = df[feature_cols]
    y = df["home_win"]

    X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, shuffle=True, random_state=42)
    model = RandomForestClassifier(n_estimators=100, random_state=42)
    model.fit(X_train, y_train)
    print("Test accuracy:", model.score(X_test, y_test))
    return model, feature_cols

# ----------------------------
# 4. Prediction Utility Module
# ----------------------------
def get_latest_stats(df, team, current_date):
    prev_games = df[((df['home_team'] == team) | (df['away_team'] == team)) & (df['date'] < current_date)].sort_values('date').tail(10)
    games_played = len(prev_games)
    wins = 0
    gf = 0
    ga = 0
    last_date = pd.Timestamp("2000-01-01")
    for _, g in prev_games.iterrows():
        if g['home_team'] == team:
            gf += g['home_score']
            ga += g['away_score']
            if g['home_score'] > g['away_score']:
                wins += 1
            last_date = g['date']
        else:
            gf += g['away_score']
            ga += g['home_score']
            if g['away_score'] > g['home_score']:
                wins += 1
            last_date = g['date']
    win_rate = wins / games_played if games_played > 0 else 0.5
    goals_for = gf / games_played if games_played > 0 else 1.0
    goals_against = ga / games_played if games_played > 0 else 1.0
    days_since = (pd.Timestamp.today() - last_date).days
    return win_rate, goals_for, goals_against, days_since

def predict_match(model, feature_cols, df, home_team, away_team, neutral=0, tournament="Friendly"):
    today = pd.Timestamp.today()
    home_stats = get_latest_stats(df, home_team, today)
    away_stats = get_latest_stats(df, away_team, today)
    input_row = {
        "home_win_rate": home_stats[0],
        "away_win_rate": away_stats[0],
        "home_goals_for": home_stats[1],
        "home_goals_against": home_stats[2],
        "away_goals_for": away_stats[1],
        "away_goals_against": away_stats[2],
        "home_days_since": home_stats[3],
        "away_days_since": away_stats[3],
        "neutral": neutral
    }
    # One-hot tournaments
    for col in feature_cols:
        if col.startswith("tournament_"):
            input_row[col] = 1 if col == f"tournament_{tournament}" else 0

    # Ensure all columns exist
    for col in feature_cols:
        if col not in input_row:
            input_row[col] = 0
    X_pred = pd.DataFrame([input_row])[feature_cols]
    proba = model.predict_proba(X_pred)[0][1]
    return proba

# ----------------------------
# 5. Main Routine
# ----------------------------
if __name__ == "__main__":
    df = load_and_prepare_data(TRAINING_FILE)
    df = add_features(df)
    model, feature_cols = train_model(df)
    joblib.dump((model, feature_cols), "football_win_predictor.joblib")

    # EXAMPLE PREDICTION
    home_team = "China PR"
    away_team = "Indonesia"
    probability = predict_match(model, feature_cols, df, home_team, away_team, neutral=0, tournament="Friendly")
    print(f"Predicted probability that {home_team} (home) beats {away_team}: {probability:.2%}")