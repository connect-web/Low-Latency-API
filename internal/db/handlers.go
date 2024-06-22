package db

import (
	"database/sql"
	"encoding/json"
	"github.com/connect-web/Low-Latency-API/internal/model"
)

// HandleSimplePlayerRow processes a SQL row into a SimplePlayer struct.
func HandleSimplePlayerRow(rows *sql.Rows) (model.SimplePlayer, error) {
	var plr model.SimplePlayer
	var minigameBytes, skillRatioBytes []byte

	if err := rows.Scan(&plr.Username, &skillRatioBytes, &minigameBytes); err != nil {
		return plr, err
	}

	if err := json.Unmarshal(minigameBytes, &plr.Minigames); err != nil {
		return plr, err
	}
	if err := json.Unmarshal(skillRatioBytes, &plr.SkillRatios); err != nil {
		return plr, err
	}

	return plr, nil
}
