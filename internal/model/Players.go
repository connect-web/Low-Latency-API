package model

type SimplePlayer struct {
	Username        string
	PID             int
	CombatLevel     int
	TotalLevel      int
	TotalExperience int64
	Skills          map[string]int // Skill : Experience directly from Hiscores
	SkillLevels     map[string]int
	SkillRatios     map[string]float32 // 32 bits will have enough useful data
	Minigames       map[string]int     // Minigame/activity : score directly from Hiscores
}

type Player struct {
	Username string

	SkillGains      map[string]float64 // these are daily stats so float is required.
	SkillGainsRatio map[string]float64
	MinigameGains   map[string]float64

	Skills          map[string]int64 // Skill : Experience directly from Hiscores
	SkillLevels     map[string]int
	SkillRatios     map[string]float64 // 32 bits will have enough useful data
	Minigames       map[string]int
	CombatLevel     int
	TotalLevel      int
	TotalExperience int64
}
