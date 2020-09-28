package logger

import (
	"time"
	"fmt"
)

func format(log string) string {
	return fmt.Sprintf("%v: %v", time.Now().Format("2006-01-02 15:04:05"), log)
}

// Log logs the passed values
func Log(i string) (log string) {
	return ""
}
