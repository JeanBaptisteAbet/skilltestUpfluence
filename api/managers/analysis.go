package managers

import (
	"fmt"
	"skillsTestUpfluence/models"
	"time"
)

func Analysis(duration time.Duration, dimension string) (models.Analysis, error) {

	//create new analysis
	analysis := models.NewAnalysis(duration, dimension)

	// start the analysis
	err := analysis.ExecuteAlysis()
	if err != nil {
		return models.Analysis{}, fmt.Errorf("Function: Analysis (managers). Error’s condition: Executing analysis. Error: %v", err)
	}

	return analysis, nil
}
