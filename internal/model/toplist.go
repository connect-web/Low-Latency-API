package model

type SkillToplist struct {
	Id        int      `json:"id"`
	Skills    []string `json:"skills"`
	Count     int      `json:"count"`
	PlayerIds []int    `json:"playerIds"`
}

type MinigameToplist struct {
	Minigame string                 `json:"minigame"`
	Count    int                    `json:"count"`
	Metrics  MachineLearningMetrics `json:"Metrics"`
}

type MachineLearningMetrics struct {
	ROCAUC         float64 `json:"ROC-AUC"`
	MeanAccuracy   float64 `json:"Mean Accuracy"`
	RecallClass0   float64 `json:"Recall Class 0"`
	RecallClass1   float64 `json:"Recall Class 1"`
	AccuracyClass0 float64 `json:"Accuracy Class 0"`
	AccuracyClass1 float64 `json:"Accuracy Class 1"`
}
