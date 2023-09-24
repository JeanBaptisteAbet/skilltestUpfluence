package models

import (
	"bufio"
	"fmt"
	"net/http"
	"skillsTestUpfluence/constant"
	"time"
)

type SSEClient struct{}

func (c SSEClient) Stream(address string, duration time.Duration) (chan Event, error) {

	eventChan := make(chan Event)
	req, err := http.NewRequest("GET", address, nil)
	if err != nil {
		return nil, fmt.Errorf("Function: Stream (models). Error’s condition: Creating request. Error: %v", err)
	}

	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Connection", "keep-alive")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Function: Stream (models). Error’s condition: Executing request. Error: %v", err)
	}

	go func() {
		continueLoop := true
		time.AfterFunc(duration, func() {
			close(eventChan)
			continueLoop = false
		})

		scanner := bufio.NewScanner(resp.Body)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			eventString := scanner.Text()
			if !continueLoop { //Security before sending in a closed channel
				return
			}
			if len(eventString) == 0 {
				continue
			}
			event, err := NewEvent(eventString)
			if err != nil {
				GetLogger().SendLog(constant.LOG_LEVEL_ERROR, fmt.Sprintf("Function: Stream (controllers). Error’s condition: Creating event from '%v'. Error: %v", eventString, err))
				continue
			}
			eventChan <- event
		}
	}()

	return eventChan, nil
}
