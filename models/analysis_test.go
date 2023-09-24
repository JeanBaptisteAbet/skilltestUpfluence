package models

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func init() {
	os.Setenv("ENVIRONMENT", "test")
}

func TestAnalysis(t *testing.T) {
	analysis := NewAnalysis(5*time.Second, "likes")

	analysisExepected := Analysis{
		TotalPost:        19, //5*4-1
		MinimumTimeStamp: 1679835328,
		MaximumTimeStamp: 1695527768,
		P50:              80,
		P90:              4738,
		P99:              4738,
		dimension:        "likes",
		duration:         5 * time.Second,
	}
	err := analysis.ExecuteAlysis()
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if analysisExepected != analysis {
		t.Errorf(" analysis is different from analysisExepected ")
	}
	fmt.Println(analysis)
}
