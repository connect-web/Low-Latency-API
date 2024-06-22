package model

type SimplePlayer struct {
	Username    string
	PID         int
	Skills      map[string]int // Skill : Experience directly from Hiscores
	Minigames   map[string]int // Minigame/activity : score directly from Hiscores
	SkillLevels map[string]int
	SkillRatios map[string]float32 // 32 bits will have enough useful data
}
