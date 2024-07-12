package model

type ProfileStats struct {
	BotsTracked      int
	BotsBanned       int
	BannedExperience int64
	BotsAddedToday   int
}

var (
	TotalBans = 0                      // ON STARTUP WE FETCH THIS AND USE IT GLOBALLY
	UserStats = map[int]ProfileStats{} // This will be updated every 12 / 24 Hours
)

func UpdateGlobalStatistics() {
	// pass
}
