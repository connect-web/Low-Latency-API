package loadrow

import (
	"database/sql"
	"fmt"
	"github.com/connect-web/Low-Latency-API/internal/model"
	"github.com/connect-web/Low-Latency-API/internal/util"
	"strconv"
)

func Player(rows *sql.Rows) ([]model.Player, error) {
	var results []model.Player

	for rows.Next() {
		entry := model.Player{}
		var skills, skillRatios, combatLevel, TotalExperience, TotalLevel, skillLevels, minigames, skillGains, skillGainsRatio, minigameGains []byte

		scanErr := rows.Scan(
			&entry.Username,
			&combatLevel, &TotalExperience, &TotalLevel,
			&skills, &skillRatios, &skillLevels, &minigames,
			&skillGains, &skillGainsRatio, &minigameGains,
		)
		if scanErr != nil {
			fmt.Printf("Player scan: %s\n", scanErr.Error())
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

		if combatLevel == nil {
			entry.CombatLevel = 3
		} else {
			byteToInt, err := strconv.Atoi(string(combatLevel))
			if err == nil {
				entry.CombatLevel = byteToInt
			} else {
				fmt.Printf("Failed to convert %s into int\n", string(combatLevel))
			}
		}

		entry.TotalExperience = 0
		byteToIntExp, err := strconv.ParseInt(string(TotalExperience), 10, 64)
		if err == nil {
			entry.TotalExperience = byteToIntExp
		}

		byteToIntTotalLvl, err := strconv.Atoi(string(TotalLevel))
		if err == nil {
			entry.TotalLevel = byteToIntTotalLvl
		}

		results = append(results, entry)
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("Player: %s\n", err.Error())
		return nil, err
	}

	return results, nil
}
