package Scanner

import (
	"database/sql"
	"encoding/json"
	"github.com/connect-web/Low-Latency-API/internal/model"
)

func ScanPlayerRows(rows *sql.Rows) ([]model.Player, error) {
	var results []model.Player

	for rows.Next() {
		entry := model.Player{}
		var skills, skillRatios, skillLevels, minigames, skillGains, skillGainsRatio, minigameGains []byte

		scanErr := rows.Scan(
			&entry.Username,
			&entry.CombatLevel, &entry.TotalExperience, &entry.TotalLevel,
			&skills, &skillRatios, &skillLevels, &minigames,
			&skillGains, &skillGainsRatio, &minigameGains,
		)
		if scanErr != nil {
			return nil, scanErr
		}

		// now convert bytes into maps

		entry.Skills = decodeJSONToInt64Map(skills)
		entry.SkillRatios = decodeJSONToFloat64Map(skillRatios)
		entry.SkillLevels = decodeJSONToIntMap(skillLevels)
		entry.Minigames = decodeJSONToIntMap(minigames)
		entry.SkillGains = decodeJSONToFloat64Map(skillGains)
		entry.SkillGainsRatio = decodeJSONToFloat64Map(skillGainsRatio)
		entry.MinigameGains = decodeJSONToFloat64Map(minigameGains)

		results = append(results, entry)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func decodeJSONToInt64Map(data []byte) map[string]int64 {
	var result map[string]int64
	if err := json.Unmarshal(data, &result); err != nil {
		return make(map[string]int64)
	}
	return result
}

func decodeJSONToIntMap(data []byte) map[string]int {
	var result map[string]int
	if err := json.Unmarshal(data, &result); err != nil {
		return make(map[string]int)
	}
	return result
}

func decodeJSONToFloat64Map(data []byte) map[string]float64 {
	var result map[string]float64
	if err := json.Unmarshal(data, &result); err != nil {
		return make(map[string]float64)
	}
	return result
}
