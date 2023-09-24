package main

import (
	"io"
	"net/http"
	"os"
	"testing"
)

func init() {
	os.Setenv("ENVIRONMENT", "test")
}

func TestRouteAnalysis(t *testing.T) {
	go Startserver()
	client := &http.Client{}
	type expextedResult struct {
		Body   string
		Status int
	}
	mapRequestResult := make(map[*http.Request]expextedResult)

	// ---------  write all request to test --------- //
	// valid request with time in Second
	requestSecond, _ := http.NewRequest("GET", "http://localhost:8080/analysis", nil)
	query := requestSecond.URL.Query()
	query.Add("dimension", "likes")
	query.Add("duration", "5s")
	requestSecond.URL.RawQuery = query.Encode()
	mapRequestResult[requestSecond] = expextedResult{
		Status: 200,
		Body:   `{"total_post":19,"minimum_timestamp":1679835328,"maximum_timestamp":1695527768,"p50":80,"p90":4738,"p99":4738}`,
	}

	// valid request with time in Minute
	requestMinute, _ := http.NewRequest("GET", "http://localhost:8080/analysis", nil)
	query = requestMinute.URL.Query()
	query.Add("dimension", "likes")
	query.Add("duration", "1m")
	requestMinute.URL.RawQuery = query.Encode()
	mapRequestResult[requestMinute] = expextedResult{
		Status: 200,
		Body:   `{"total_post":239,"minimum_timestamp":1679835328,"maximum_timestamp":1695527768,"p50":80,"p90":4738,"p99":4738}`,
	}

	// // Commented because 1h for a test is a bit too long
	// // valid request with time in Hours
	// requestHours, _ := http.NewRequest("GET", "http://localhost:8080/analysis", nil)
	// query = requestHours.URL.Query()
	// query.Add("dimension", "likes")
	// query.Add("duration", "1h")
	// requestHours.URL.RawQuery = query.Encode()
	// mapRequestResult[requestHours] = expextedResult{
	// 	Status: 200,
	// }

	// request wrong route
	requestWrongRoute, _ := http.NewRequest("GET", "http://localhost:8080/wrongroute", nil)
	mapRequestResult[requestWrongRoute] = expextedResult{
		Status: 404,
	}
	// request wrong Methode
	requestWrongMethode, _ := http.NewRequest("POST", "http://localhost:8080/analysis", nil)
	mapRequestResult[requestWrongMethode] = expextedResult{
		Status: 404,
	}

	// ---------  test all request --------- //
	for request, expextedResult := range mapRequestResult {
		resp, err := client.Do(request)
		if err != nil {
			t.Errorf("unexpected error %v", err)
			continue
		}
		if expextedResult.Status != resp.StatusCode {
			t.Errorf("Wrong status code for request '%v'. Got '%v'", request, resp.StatusCode)
		}

		if expextedResult.Status == 200 {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("unexpected error %v", err)
			}
			if string(body) != expextedResult.Body {
				t.Errorf("Wrong body for request '%v'. Got '%v'", request, string(body))
			}
		}
	}
}
