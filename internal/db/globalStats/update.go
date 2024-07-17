package globalStats

import (
	"errors"
	"fmt"
	"github.com/connect-web/Low-Latency-API/internal/db"
	"github.com/connect-web/Low-Latency-API/internal/model"
	"log"
	"time"
)

func GlobalStatsWorker() {
	for {
		UpdateGlobalStatistics()
		time.Sleep(1 * time.Hour)
	}
}

func UpdateGlobalStatistics() {
	if !shouldUpdateGlobalStatistics() {
		return
	}
	ban_count, err := queryTotalBans()
	if err != nil {
		return
	}
	model.Last_updated = time.Now()
	model.TotalBans = ban_count
}

func queryTotalBans() (int, error) {
	var total_bans int

	query := "SELECT BANS FROM PROFILES.global_stats"
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

	row := client.DB.QueryRow(query)

	scanErr := row.Scan(&total_bans)

	return total_bans, scanErr
}

func shouldUpdateGlobalStatistics() bool {
	duration := time.Now().Sub(model.Last_updated)
	hoursSinceLastUpdate := duration.Hours()
	fmt.Printf("Hours since last update: %.2f\n", hoursSinceLastUpdate)

	return 12 < hoursSinceLastUpdate
	// pass
}
