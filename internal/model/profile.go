package model

import "time"

type ProfileStats struct {
	Username         string
	BotsTracked      int
	BotsBanned       int
	BotsPlayerIds    []int
	BannedExperience int64
	PlayersAdded     int
}

var (
	LowLatencyStatistics = GlobalStatistics{Last_updated: time.Date(1990, 1, 1, 1, 1, 1, 1, time.UTC)}
)

type GlobalStatistics struct {
	Bans            int
	Skills          map[string]int64 // Skill : Experience directly from Hiscores
	Minigames       map[string]int64 // Minigame/activity : score directly from Hiscores
	TotalExperience int64
	SuspiciousUsers int
	Last_updated    time.Time
}
