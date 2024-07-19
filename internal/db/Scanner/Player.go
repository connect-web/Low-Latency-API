package Scanner

import (
	"database/sql"
	"github.com/connect-web/Low-Latency-API/internal/model"
	"github.com/connect-web/Low-Latency-API/internal/util"
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

		entry.Skills = util.DecodeJSONToInt64Map(skills)
		entry.SkillRatios = util.DecodeJSONToFloat64Map(skillRatios)
		entry.SkillLevels = util.DecodeJSONToIntMap(skillLevels)
		entry.Minigames = util.DecodeJSONToIntMap(minigames)
		entry.SkillGains = util.DecodeJSONToFloat64Map(skillGains)
		entry.SkillGainsRatio = util.DecodeJSONToFloat64Map(skillGainsRatio)
		entry.MinigameGains = util.DecodeJSONToFloat64Map(minigameGains)

		results = append(results, entry)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
