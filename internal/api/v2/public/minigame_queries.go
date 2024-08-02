package public

import (
	"errors"
	"github.com/connect-web/Low-Latency-API/internal/db"
	"github.com/connect-web/Low-Latency-API/internal/db/loadrow"
	"github.com/connect-web/Low-Latency-API/internal/model"
	"github.com/lib/pq"
	"log"
)

var MINIGAME_TOPLIST_QUERY = `
select 
    id, activities, amount
from grouped.minigames
ORDER BY amount DESC;
`
var MINIGAME_TOPLIST_USER_QUERY = `
SELECT
    p.name,
    COALESCE(pls.combat_level, 3), COALESCE(pls.overall, 0), COALESCE(pls.total_level, 23),
    pl.skills_experience, pl.skills_ratio, pl.skills_levels, pl.minigames,
    pg.skills_experience, pg.skills_ratio, pg.minigames

FROM player_gains pg
LEFT JOIN PLAYERS P on p.id = pg.playerid
LEFT JOIN player_live pl on pl.playerid = pg.playerid
LEFT JOIN player_live_stats pls on pls.playerid = pg.playerid
WHERE
    pg.playerid = ANY(
    select
        unnest(s.playerids)
    from grouped.minigames s
    where id = $1
    );
`

func QueryMinigameToplist() ([]model.SkillToplist, error) {
	results := []model.SkillToplist{}
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
		entry := model.SkillToplist{}
		scanErr := rows.Scan(&entry.Id, pq.Array(&entry.Skills), &entry.Count)
		if scanErr == nil {
			results = append(results, entry)
		}
	}
	return results, nil
}

func QueryMinigameToplistUsers(minigame int) ([]model.Player, error) {
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
