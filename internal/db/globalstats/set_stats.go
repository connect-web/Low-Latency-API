package globalstats

import (
	"errors"
	"fmt"
	"github.com/connect-web/Low-Latency-API/internal/db"
	"github.com/connect-web/Low-Latency-API/internal/model"
	"github.com/connect-web/Low-Latency-API/internal/util"
	"log"
	"time"
)

func GetGlobalStats() (model.GlobalStatistics, error) {
	total_stats := model.GlobalStatistics{}

	skills, minigames, err := getSumBannedStats()
	if err != nil {
		fmt.Printf("Failed to GetGlobalStats: %s\n", err.Error())
		return total_stats, err
	}
	bans, banErr := getTotalBans()
	if banErr != nil {
		fmt.Printf("Failed to GetGlobalStats: %s\n", banErr.Error())
		return total_stats, banErr
	}

	total_users, userErr := getSuspiciousUserCount()
	if userErr != nil {
		fmt.Printf("Failed to get global stats for sus users: %s\n", userErr.Error())
		return total_stats, banErr
	}

	total_stats.Bans = bans
	total_stats.Skills = skills
	total_stats.Minigames = minigames
	total_stats.TotalExperience = getTotalExperience(skills)
	total_stats.SuspiciousUsers = total_users
	total_stats.Last_updated = time.Now()

	return total_stats, nil
}

var banned_user_data_sum = `
WITH filtered_skills AS (
    SELECT
        pl.skills_experience
    FROM not_found nf
    LEFT JOIN player_live pl ON pl.playerid = nf.playerid
    WHERE jsonb_typeof(pl.skills_experience) = 'object'
),
expanded_skills AS (
    SELECT
        key AS skill_name,
        CAST(value AS numeric) AS value
    FROM filtered_skills,
    jsonb_each_text(filtered_skills.skills_experience)
),
filtered_expanded_skills AS (
    SELECT
        skill_name,
        value
    FROM expanded_skills
    WHERE value >= 0
),
aggregated_skills AS (
    SELECT
        skill_name,
        SUM(value) AS total_value
    FROM filtered_expanded_skills
    GROUP BY skill_name
),
filtered_minigames AS (
    SELECT
        pl.minigames
    FROM not_found nf
    LEFT JOIN player_live pl ON pl.playerid = nf.playerid
    WHERE jsonb_typeof(pl.minigames) = 'object'
),
expanded_minigames AS (
    SELECT
        key AS minigame_name,
        CAST(value AS numeric) AS value
    FROM filtered_minigames,
    jsonb_each_text(filtered_minigames.minigames)
),
filtered_expanded_minigames AS (
    SELECT
        minigame_name,
        value
    FROM expanded_minigames
    WHERE value >= 0
),
aggregated_minigames AS (
    SELECT
        minigame_name,
        SUM(value) AS total_value
    FROM filtered_expanded_minigames
    GROUP BY minigame_name
)
SELECT 
    (SELECT jsonb_object_agg(skill_name, total_value) FROM aggregated_skills) AS aggregated_skills_experience,
    (SELECT jsonb_object_agg(minigame_name, total_value) FROM aggregated_minigames) AS aggregated_minigames;
`

var total_banned_users = "SELECT COUNT(*) FROM not_found"

func getTotalBans() (int, error) {
	var total_bans int
	client := db.NewDBClient()
	if connectErr := client.Connect(); connectErr != nil {
		log.Println(connectErr.Error())
		return total_bans, errors.New("Internal server error")
	}
	defer func() {
		if closeConnErr := client.Close(); closeConnErr != nil {
			log.Printf("Failed to close database connection: %v", closeConnErr)
		}
	}()

	row := client.DB.QueryRow(total_banned_users)

	scanErr := row.Scan(&total_bans)

	return total_bans, scanErr
}

func getSuspiciousUserCount() (int, error) {
	var total_players int
	client := db.NewDBClient()
	if connectErr := client.Connect(); connectErr != nil {
		log.Println(connectErr.Error())
		return total_players, errors.New("Internal server error")
	}
	defer func() {
		if closeConnErr := client.Close(); closeConnErr != nil {
			log.Printf("Failed to close database connection: %v", closeConnErr)
		}
	}()

	row := client.DB.QueryRow("SELECT COUNT(*) FROM PLAYERS;")

	scanErr := row.Scan(&total_players)

	return total_players, scanErr
}

func getSumBannedStats() (map[string]int64, map[string]int64, error) {

	client := db.NewDBClient()
	if connectErr := client.Connect(); connectErr != nil {
		log.Println(connectErr.Error())
		return nil, nil, errors.New("Internal server error")
	}
	defer func() {
		if closeConnErr := client.Close(); closeConnErr != nil {
			log.Printf("Failed to close database connection: %v", closeConnErr)
		}
	}()

	row := client.DB.QueryRow(banned_user_data_sum)

	var skills, minigames []byte
	scanErr := row.Scan(&skills, &minigames)

	skill_map := util.DecodeJSONToInt64Map(skills)
	minigame_map := util.DecodeJSONToInt64Map(minigames)

	return skill_map, minigame_map, scanErr
}

func getTotalExperience(skills map[string]int64) int64 {
	var total_experience int64
	for skill, value := range skills {
		if skill == "Overall" {
			continue
		}
		total_experience += value
	}
	return total_experience
}
