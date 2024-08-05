package globalstats

import (
	"github.com/connect-web/Low-Latency-API/internal/model"
	"time"
)

func GlobalStatsWorker() {
	for {
		UpdateGlobalStatistics()
		time.Sleep(1 * time.Minute)
	}
}

func UpdateGlobalStatistics() {
	if !shouldUpdateGlobalStatistics() {
		return
	}
	global_stats, err := GetGlobalStats()
	if err == nil {
		model.LowLatencyStatistics = global_stats
	}
}

func shouldUpdateGlobalStatistics() bool {
	duration := time.Now().Sub(model.LowLatencyStatistics.Last_updated)
	hoursSinceLastUpdate := duration.Hours()
	//fmt.Printf("Hours since last update: %.2f\n", hoursSinceLastUpdate)

	return 1 < hoursSinceLastUpdate
	// pass
}
