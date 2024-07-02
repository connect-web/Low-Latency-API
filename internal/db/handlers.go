package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/connect-web/Low-Latency-API/internal/model"
	"log"
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

func HandlePlayerSkillsRow(rows *sql.Rows) (model.Player, error) {
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

// todo CREATE A LOGGING FEATURE FOR WHEN LEN(bytes) for json dicts or columns are nulled, This should respond as nulled instead of throwing errors, Also when logged can later be checked to fix null issues or find issues in Transfer Script.

func HandlePlayerRowMinigames(rows *sql.Rows) (model.Player, error) {
	var plr model.Player
	var minigameBytes, skillLevelBytes, skillExperienceGainBytes, minigameGainBytes []byte

	if err := rows.Scan(&plr.Username, &skillLevelBytes, &minigameBytes, &skillExperienceGainBytes, &minigameGainBytes); err != nil {
		fmt.Println("Error scanning user")
		return plr, err
	}

	if len(skillLevelBytes) > 0 {
		if err := json.Unmarshal(skillLevelBytes, &plr.SkillLevels); err != nil {
			log.Printf("Error unmarshalling skillLevelBytes: %v, data: %s", err, string(skillLevelBytes))
			return plr, err
		}
	}
	if len(minigameBytes) > 0 {
		if err := json.Unmarshal(minigameBytes, &plr.Minigames); err != nil {
			log.Printf("Error unmarshalling minigameBytes: %v, data: %s", err, string(minigameBytes))
			return plr, err
		}
	}
	if len(skillExperienceGainBytes) > 0 && string(skillExperienceGainBytes) != "null" {
		if err := json.Unmarshal(skillExperienceGainBytes, &plr.SkillGains); err != nil {
			log.Printf("Error unmarshalling skillExperienceGainBytes: %v, data: %s", err, string(skillExperienceGainBytes))
			return plr, err
		}
	}
	if len(minigameGainBytes) > 0 && string(minigameGainBytes) != "null" {
		if err := json.Unmarshal(minigameGainBytes, &plr.MinigameGains); err != nil {
			log.Printf("Error unmarshalling minigameGainBytes: %v, data: %s", err, string(minigameGainBytes))
			return plr, err
		}
	}

	return plr, nil
}
