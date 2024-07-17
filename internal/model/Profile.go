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
	Last_updated = time.Now().Add(-500 * time.Hour)
	TotalBans    = 0 // ON STARTUP WE FETCH THIS AND USE IT GLOBALLY
)
