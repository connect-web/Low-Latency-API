package user

import (
	"errors"
	"github.com/connect-web/Low-Latency-API/internal/db"
	"github.com/connect-web/Low-Latency-API/internal/model"
	"github.com/lib/pq"
	"log"
)

func CreateProfileStatistics(username string) error {
	query := `
	INSERT INTO profiles.user_stats (id)
	(SELECT id from users.accs where name = $1)
	ON CONFLICT DO NOTHING;
	`
	client := db.NewDBClient()
	if connectErr := client.Connect(); connectErr != nil {
		log.Println(connectErr.Error())
		return errors.New("Internal server error")
	}
	defer func() {
		if closeConnErr := client.Close(); closeConnErr != nil {
			log.Printf("Failed to close database connection: %v", closeConnErr)
		}
	}()
	_, err := client.DB.Exec(query, username)
	return err
}

// Profile stats are only created when visiting the profile page after login
// This could be used to see if accounts are being signed up by automated software that does not run the site's javascript.
func FetchOrCreateProfile(username string) (model.ProfileStats, error) {
	stats, err := GetProfileStatistics(username)
	if err != nil && err.Error() == "No stats found!" {
		createErr := CreateProfileStatistics(username)
		if createErr != nil {
			log.Println("Failed to create profile for %s due to %s\n", username, createErr.Error())
			return stats, createErr
		}
		stats, err = GetProfileStatistics(username)
	}
	return stats, err
}

func GetProfileStatistics(username string) (model.ProfileStats, error) {
	query := `
	SELECT
	    us.id,
	    us.bots_tracked,
	    us.bots_banned,
	    us.bots_playerids,
	    us.banned_experience,
	    us.players_added
	FROM users.accs acc
	LEFT JOIN profiles.user_stats us on us.id = acc.id
	where acc.name = $1
	`
	var profileStats = model.ProfileStats{}

	client := db.NewDBClient()
	if connectErr := client.Connect(); connectErr != nil {
		log.Println(connectErr.Error())
		return profileStats, errors.New("Internal server error")
	}
	defer func() {
		if closeConnErr := client.Close(); closeConnErr != nil {
			log.Printf("Failed to close database connection: %v", closeConnErr)
		}
	}()

	row := client.DB.QueryRow(query, username)
	var userId int
	scanErr := row.Scan(
		&userId,
		&profileStats.BotsTracked,
		&profileStats.BotsBanned,
		pq.Array(&profileStats.BotsPlayerIds),
		&profileStats.BannedExperience,
		&profileStats.PlayersAdded,
	)
	if userId == 0 {
		return model.ProfileStats{}, errors.New("No stats found!")
	}

	return profileStats, scanErr
}
