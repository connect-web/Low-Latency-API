package Scanner

import (
	"database/sql"
	"github.com/connect-web/Low-Latency-API/internal/model"
	"github.com/lib/pq"
)

func ScanPlayerRows(rows *sql.Rows) ([]model.Player, error) {
	var results []model.Player

	for rows.Next() {
		entry := model.Player{}
		scanErr := rows.Scan(
			&entry.Username,
			&entry.CombatLevel, &entry.TotalExperience, &entry.TotalLevel,
			pq.Array(&entry.Skills), pq.Array(&entry.SkillRatios), pq.Array(&entry.SkillLevels), pq.Array(&entry.Minigames),
			pq.Array(&entry.SkillGains), pq.Array(&entry.SkillGainsRatio), pq.Array(&entry.MinigameGains),
		)
		if scanErr != nil {
			return nil, scanErr
		}
		results = append(results, entry)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
