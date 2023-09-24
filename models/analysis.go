package models

import (
	"time"
)

type Analysis struct {
	TotalPost        int   `json:"total_post"`
	MinimumTimeStamp int64 `json:"minimum_timestamp"`
	MaximumTimeStamp int64 `json:"maximum_timestamp"`
	P50              int   `json:"p50"`
	P90              int   `json:"p90"`
	P99              int   `json:"p99"`

	duration  time.Duration
	dimension string
}

func NewAnalysis(duration time.Duration, dimension string) Analysis {
	return Analysis{
		duration:  duration,
		dimension: dimension,
	}
}
