package models

import (
	"fmt"
	"os"
	"slices"
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

// ExecuteAlysis will complete Analysis information
// will return an error if there is one or nil if the analysis completed corectly
func (a *Analysis) ExecuteAlysis() error {
	sseClient := GetNewClient()

	eventChan, err := sseClient.Stream(os.Getenv("STREAM_ADRESS"), a.duration)
	if err != nil {
		return fmt.Errorf("Function: ExecuteAlysis (models). Error’s condition: Getting stream of event. Error: %v", err)
	}
	sliceDimension := []int{}

	for event := range eventChan {

		a.TotalPost++

		if a.MinimumTimeStamp == 0 || event.TimeStamp < a.MinimumTimeStamp {
			a.MinimumTimeStamp = event.TimeStamp
		}
		if event.TimeStamp > a.MaximumTimeStamp {
			a.MaximumTimeStamp = event.TimeStamp
		}

		switch a.dimension {
		case "likes":
			sliceDimension = append(sliceDimension, event.Likes)
		case "comments":
			sliceDimension = append(sliceDimension, event.Comments)
		case "favorites":
			sliceDimension = append(sliceDimension, event.Favorites)
		case "retweets":
			sliceDimension = append(sliceDimension, event.Retweets)
		}

	}

	if len(sliceDimension) != a.TotalPost {
		return fmt.Errorf("Function: ExecuteAlysis (models). Error’s condition: Checking sliceDimension length. Error: Length is '%v', should be '%v'", len(sliceDimension), a.TotalPost)
	}

	slices.Sort(sliceDimension)

	p50Position := a.TotalPost >> 1 // position divideed by two
	p90Position := a.TotalPost * 90 / 100
	p99Position := a.TotalPost * 99 / 100

	a.P50 = sliceDimension[p50Position]
	a.P90 = sliceDimension[p90Position]
	a.P99 = sliceDimension[p99Position]

	return nil
}
