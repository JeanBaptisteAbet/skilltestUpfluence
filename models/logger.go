package models

import (
	"fmt"
	"time"
)

// basic logger
type Logger struct{}

func GetLogger() Logger {
	return Logger{}
}

func (l Logger) SendLog(level, message string) {
	fmt.Println(fmt.Sprintf("%v:%v:%v", time.Now(), level, message))
}
