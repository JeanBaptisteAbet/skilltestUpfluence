package models

import (
	"encoding/json"
	"fmt"
)

type Event struct {
	SocialMedia string
	Likes       int
	Comments    int
	Favorites   int
	Retweets    int
	TimeStamp   int64
}

// NewEvent transform SSE data into a struct event
func NewEvent(eventString string) (Event, error) {

	eventByte := []byte(eventString)

	var mapEventBySocialMedia map[string]Event

	if err := json.Unmarshal(eventByte[6:], &mapEventBySocialMedia); err != nil {
		return Event{}, fmt.Errorf("Function: NewEvent (models). Error’s condition: Unmarshaling event '%v'. Error: %v", eventString, err)
	}
	for socialMedia, event := range mapEventBySocialMedia {
		return Event{
			SocialMedia: socialMedia,
			Likes:       event.Likes,
			Comments:    event.Comments,
			Favorites:   event.Favorites,
			Retweets:    event.Retweets,
			TimeStamp:   event.TimeStamp,
		}, nil
	}
	return Event{}, nil
}
