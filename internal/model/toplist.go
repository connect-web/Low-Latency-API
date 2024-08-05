package model

import "time"

type Toplist struct {
	Id          int       `json:"id"`
	LastUpdated time.Time `json:"lastUpdated"`
	Activities  []string  `json:"Activities"`

	LifetimeCount int `json:"lifetimeCount"`
	UnbannedCount int `json:"UnbannedCount"`
	BannedCount   int `json:"BannedCount"`

	UnbannedPlayerIds []int `json:"unbannedPlayerIds"`
	BannedPlayerIds   []int `json:"bannedPlayerIds"`
}

type MinigameToplist struct {
	Minigame    string                 `json:"minigame"`
	LastUpdated time.Time              `json:"lastUpdated"`
	Count       int                    `json:"count"`
	Metrics     MachineLearningMetrics `json:"Metrics"`
}

type MachineLearningMetrics struct {
	ROCAUC         float64 `json:"ROC-AUC"`
	MeanAccuracy   float64 `json:"Mean Accuracy"`
	RecallClass0   float64 `json:"Recall Class 0"`
	RecallClass1   float64 `json:"Recall Class 1"`
	AccuracyClass0 float64 `json:"Accuracy Class 0"`
	AccuracyClass1 float64 `json:"Accuracy Class 1"`
}
