from flask import Flask, request, jsonify
import pandas as pd
import joblib,os
from matchPredictionTraining import predict_match,TRAINING_FILE,FOLDER
import json
from pathlib import Path

# ---- Load model and features at startup ----
model, feature_cols = joblib.load(os.path.join(FOLDER,"football_win_predictor.joblib")) 

df = pd.read_csv(TRAINING_FILE)
df['date'] = pd.to_datetime(df['date'])


# ----- Flask app -----
app = Flask(__name__)

@app.route('/odds', methods=['POST'])
def calculate_odds():
    data = request.json
    
    event = data.get('event', '')
    if "vs" in event:
        home_team, away_team = [s.strip() for s in event.split("vs")]
    else:
        return jsonify({'yes': 50, 'no': 50})
    
    # Use model to predict probability
    prob = predict_match(model, feature_cols, df, home_team, away_team, neutral=0, tournament="Friendly")
    yes_odds = int(prob * 100)
    no_odds = 100 - yes_odds

    # # Get real time odds from agent
    yes_odds_rt, no_odds_rt = read_real_time_odds('latest_country_competitions.json',home_team, away_team)

    # # Combine the Machine learning model prediction and Real time odds
    ml_odds_weight = 0.8
    rt_odds_weight = 0.2
    if yes_odds_rt + no_odds_rt == 100:
        yes_odds = round(yes_odds * ml_odds_weight + yes_odds_rt * rt_odds_weight)
    no_odds = 100 - yes_odds

    return jsonify({'yes': yes_odds, 'no': no_odds})


def read_real_time_odds(file, home_team, away_team):
    base_dir = Path(__file__).resolve().parent
    file_path = base_dir / "outputs" / file

    home_key = f"{home_team.lower().replace(' ', '_')}_win"
    draw_key = 'draw_win'
    try:
        with open(file_path, "r", encoding="utf-8") as f:
            data = json.load(f)

        for entry in data:
            match = entry.get("match", "").lower().replace(".", "").strip()
            if home_team.lower() in match and away_team.lower() in match:
                predictions = entry.get("polymarket_prediction", {})
                if home_key in predictions:
                    win = round((predictions[home_key] / (1 - predictions[draw_key])) * 100)
                    if win >= 100:
                        win = 100
                    lose = 100 - win
                    return win, lose
        return 0, 0

    except (FileNotFoundError, json.JSONDecodeError) as e:
        print('error: ',e)
        return 0, 0

if __name__ == '__main__':
    app.run(port=5001)
