package model

type SkillToplist struct {
	Id        int      `json:"id"`
	Skills    []string `json:"skills"`
	Count     int      `json:"count"`
	PlayerIds []int    `json:"playerIds"`
}

type MinigameToplist struct {
	Minigame string `json:"minigame"`
	Count    int    `json:"count"`
}
