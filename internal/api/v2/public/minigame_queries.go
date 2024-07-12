package public

import (
	"errors"
	"github.com/connect-web/Low-Latency-API/internal/db"
	"github.com/connect-web/Low-Latency-API/internal/db/Scanner"
	"github.com/connect-web/Low-Latency-API/internal/model"
	"log"
)

var MINIGAME_TOPLIST_QUERY = `
SELECT 
	minigame, 
	COUNT(DISTINCT id) AS count
FROM stats.minigame_links
GROUP BY minigame
ORDER BY count DESC;
`

var MINIGAME_TOPLIST_USER_QUERY = `
	SELECT
		p.name,
		pls.combat_level, pls.overall, pls.total_level,
		pl.skills_experience, pl.skills_ratio, pl.skills_levels, pl.minigames,
		pg.skills_experience, pg.skills_ratio, pg.minigames
	FROM stats.minigame_links links
	LEFT JOIN PLAYERS P on p.id = links.id
	LEFT JOIN player_live pl on pl.playerid = links.id
	LEFT JOIN player_live_stats pls on pls.playerid = links.id
	LEFT JOIN player_gains pg on pg.playerid = links.id
	WHERE 
	    links.MINIGAME = $1
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
		if err := rows.Scan(&row.Minigame, &row.Count); err == nil {
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

	results, err := Scanner.ScanPlayerRows(rows)
	return results, err
}
