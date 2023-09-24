package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"skillsTestUpfluence/api/managers"
	"skillsTestUpfluence/constant"
	"skillsTestUpfluence/models"
	"time"
)

func Analysis(w http.ResponseWriter, r *http.Request) {

	models.GetLogger().SendLog(constant.LOG_LEVEL_INFO, fmt.Sprintf("Function: Analysis (controllers). Info: Starting request '%v'", r))

	validDimension := map[string]bool{
		"likes":     true,
		"comments":  true,
		"favorites": true,
		"retweets":  true,
	}

	// fetch param from the request
	durationString, dimensionString := r.URL.Query().Get("duration"), r.URL.Query().Get("dimension")

	//check if duration is valid and convert it to time.Duration
	duration, err := time.ParseDuration(durationString)
	if err != nil {
		models.GetLogger().SendLog(constant.LOG_LEVEL_ERROR, fmt.Sprintf("Function: Analysis (controllers). Error’s condition: Converting duration from request to time.Duration. Error: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//check if dimension is valid
	if ok := validDimension[dimensionString]; !ok {
		models.GetLogger().SendLog(constant.LOG_LEVEL_WARNING, fmt.Sprintf("Function: Analysis (controllers). Error’s condition: Checking dimensions requested. Error: Dimension '%v' unknown", dimensionString))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//call manager to execute job logic
	analysisResult, err := managers.Analysis(duration, dimensionString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		models.GetLogger().SendLog(constant.LOG_LEVEL_ERROR, fmt.Sprintf("Function: Analysis (controllers). Error’s condition: Getting analysis. Error: %v", err))
		return
	}

	//transform analysis into json
	analysisMarshaled, err := json.Marshal(analysisResult)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		models.GetLogger().SendLog(constant.LOG_LEVEL_ERROR, fmt.Sprintf("Function: Analysis (controllers). Error’s condition: Marshalling analysisResult. Error: %v", err))
		return
	}

	//return result
	w.WriteHeader(http.StatusOK)
	w.Write(analysisMarshaled)

	models.GetLogger().SendLog(constant.LOG_LEVEL_INFO, fmt.Sprintf("Function: Analysis (controllers). Info: Request '%v' finished with result '%v'. ", r, analysisResult))

}
