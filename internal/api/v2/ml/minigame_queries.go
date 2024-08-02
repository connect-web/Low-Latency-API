package ml

import (
	"encoding/json"
	"errors"
	"github.com/connect-web/Low-Latency-API/internal/db"
	"github.com/connect-web/Low-Latency-API/internal/db/loadrow"
	"github.com/connect-web/Low-Latency-API/internal/model"
	"log"
)

var MINIGAME_TOPLIST_QUERY = `
SELECT r.activity, count(distinct r.playerid), m.metrics
FROM ML.results r
LEFT JOIN ml.metrics m on m.activity = r.activity
WHERE r.activitytype = 'minigames'
GROUP BY r.activity, m.metrics
ORDER BY count(distinct r.playerid) desc;
`

var MINIGAME_TOPLIST_USER_QUERY = `
	SELECT p.name,
       pls.combat_level,
       pls.overall,
       pls.total_level,
       pl.skills_experience,
       pl.skills_ratio,
       pl.skills_levels,
       pl.minigames,
       pg.skills_experience,
       pg.skills_ratio,
       pg.minigames
FROM ML.results links
         LEFT JOIN PLAYERS P on p.id = links.playerid
         LEFT JOIN player_live pl on pl.playerid = links.playerid
         LEFT JOIN player_live_stats pls on pls.playerid = links.playerid
         LEFT JOIN player_gains pg on pg.playerid = links.playerid
WHERE links.activity = $1
AND links.activityType = 'minigames'
ORDER BY links DESC;
	`

func QueryMinigameToplist() ([]model.MinigameToplist, error) {
	results := []model.MinigameToplist{}
	client := db.NewDBClient()
	if connectErr := client.Connect(); connectErr != nil {
		log.Println(connectErr.Error())
		return results, errors.New("Internal server error")
	}
	defer func() {
		if closeConnErr := client.Close(); closeConnErr != nil {
			log.Printf("Failed to close database connection: %v", closeConnErr)
		}
	}()

	rows, queryErr := client.DB.Query(MINIGAME_TOPLIST_QUERY)
	if queryErr != nil {
		log.Println("Error executing query:", queryErr.Error())
		return results, errors.New("Internal server error")
	}
	defer rows.Close()

	for rows.Next() {
		row := model.MinigameToplist{}
		var metrics []byte
		if err := rows.Scan(&row.Minigame, &row.Count, &metrics); err == nil {
			metricErr := json.Unmarshal(metrics, &row.Metrics)
			if metricErr != nil {
				log.Printf("Failed to unmarshal metrics: %s\n", metricErr.Error())
			}
			results = append(results, row)
		}
	}
	return results, nil
}

func QueryMinigameToplistUsers(minigame string) ([]model.Player, error) {
	results := []model.Player{}

	client := db.NewDBClient()
	if connectErr := client.Connect(); connectErr != nil {
		log.Println(connectErr.Error())
		return results, errors.New("Internal server error")
	}
	defer func() {
		if closeConnErr := client.Close(); closeConnErr != nil {
			log.Printf("Failed to close database connection: %v", closeConnErr)
		}
	}()

	rows, queryErr := client.DB.Query(MINIGAME_TOPLIST_USER_QUERY, minigame)
	if queryErr != nil {
		log.Println("Error executing query:", queryErr.Error())
		return results, errors.New("Internal server error")
	}
	defer rows.Close()
	results, err := loadrow.Player(rows)
	return results, err
}
