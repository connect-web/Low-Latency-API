package public

import (
	"errors"
	"github.com/connect-web/Low-Latency-API/internal/db"
	"github.com/connect-web/Low-Latency-API/internal/db/Scanner"
	"github.com/connect-web/Low-Latency-API/internal/model"
	"github.com/lib/pq"
	"log"
)

var SKILLER_TOPLIST_QUERY = `
select 
    s.id, s.skills, sc.amount
from grouped.skills_count sc
	LEFT JOIN grouped.skills s on s.id = sc.id
	LEFT JOIN GROUPED.skills_users su on su.id = sc.id
ORDER BY SC.amount DESC
LIMIT 100;
`

var SKILLER_TOPLIST_USERS_QUERY = `
SELECT
    p.name,
    pls.combat_level, pls.overall, pls.total_level,
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
    from grouped.skills_users s
    where id = $1
    );
`

func QuerySkillToplist() ([]model.SkillToplist, error) {
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

	rows, queryErr := client.DB.Query(SKILLER_TOPLIST_QUERY)
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

func QuerySkillToplistUsers(skillId int) ([]model.Player, error) {
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

	rows, queryErr := client.DB.Query(SKILLER_TOPLIST_USERS_QUERY, skillId)
	if queryErr != nil {
		log.Println("Error executing query:", queryErr.Error())
		return results, errors.New("Internal server error")
	}
	defer rows.Close()

	results, err := Scanner.ScanPlayerRows(rows)

	return results, err
}
