package db

import (
	"database/sql"
	"encoding/json"
	"github.com/connect-web/Low-Latency-API/internal/model"
)

func HandleSimplePlayerRowRatio(rows *sql.Rows) (model.SimplePlayer, error) {
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

func HandleSimplePlayerRowExperience(rows *sql.Rows) (model.SimplePlayer, error) {
	var plr model.SimplePlayer
	var minigameBytes, skillBytes []byte

	if err := rows.Scan(&plr.Username, &skillBytes, &minigameBytes); err != nil {
		return plr, err
	}

	if err := json.Unmarshal(minigameBytes, &plr.Minigames); err != nil {
		return plr, err
	}
	if err := json.Unmarshal(skillBytes, &plr.Skills); err != nil {
		return plr, err
	}

	return plr, nil
}

func HandleSimplePlayerRowLevel(rows *sql.Rows) (model.SimplePlayer, error) {
	var plr model.SimplePlayer
	var minigameBytes, skillLevelBytes []byte

	if err := rows.Scan(&plr.Username, &skillLevelBytes, &minigameBytes); err != nil {
		return plr, err
	}

	if err := json.Unmarshal(minigameBytes, &plr.Minigames); err != nil {
		return plr, err
	}
	if err := json.Unmarshal(skillLevelBytes, &plr.SkillLevels); err != nil {
		return plr, err
	}

	return plr, nil
}

func HandlePlayerRow(rows *sql.Rows) (model.Player, error) {
	var plr model.Player
	var minigameBytes, skillLevelBytes, skillExperienceGainBytes []byte

	if err := rows.Scan(&plr.Username, &skillLevelBytes, &minigameBytes, &skillExperienceGainBytes); err != nil {
		return plr, err
	}

	if err := json.Unmarshal(minigameBytes, &plr.Minigames); err != nil {
		return plr, err
	}
	if err := json.Unmarshal(skillLevelBytes, &plr.SkillLevels); err != nil {
		return plr, err
	}
	if err := json.Unmarshal(skillExperienceGainBytes, &plr.SkillGains); err != nil {
		return plr, err
	}

	return plr, nil
}
