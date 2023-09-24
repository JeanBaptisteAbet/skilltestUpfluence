package models

import (
	"testing"
)

func TestNewEvent(t *testing.T) {
	eventFromStream := `data: {"pin":{"id":100868824,"title":"","description":" ","links":"https://www.savoringthegood.com/masterbuilt-smoker-recipes/","likes":1,"comments":2,"saves":0,"repins":0,"timestamp":1694530994,"post_id":"272327108711925220"}}`
	eventExepected := Event{
		SocialMedia: "pin",
		Likes:       1,
		Retweets:    0,
		Comments:    2,
		Favorites:   0,
		TimeStamp:   1694530994,
	}
	event, err := NewEvent(eventFromStream)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if event != eventExepected {
		t.Errorf("Event '%v' is different from eventExpected '%v' ", event, eventExepected)
	}
}
