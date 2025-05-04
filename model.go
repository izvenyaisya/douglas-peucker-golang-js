package main

import "time"

type Point struct {
	Lat       float64   `json:"lat"`
	Lng       float64   `json:"lng"`
	Altitude  float64   `json:"altitude"`
	Speed     float64   `json:"speed"`
	Timestamp time.Time `json:"timestamp"`
	Paused    bool      `json:"paused"`
}

type AnalyzeRequest struct {
	UserID int32  `json:"userId"`
	RawPoints []Point `json:"rawPoints"`
}

type AnalyzeResponse struct {
	UserID       int32  `json:"userId"`
	StartAt      time.Time `json:"startAt"`
	EndAt        time.Time `json:"endAt"`
	Distance     float64 `json:"distance"`
	Calories     float64 `json:"calories"`
	AverageSpeed float64 `json:"averageSpeed"`
	Elevation    float64 `json:"elevation"`
	Points       []Point `json:"points"`
}
