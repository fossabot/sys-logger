package logger

import (
	"time"
	"fmt"
)

// Log logs the passed values
func Log(i string) (log string) {
	log = fmt.Sprintf("%v: %v", time.Now().Format("2006-01-02 15:04:05"), i)
	fmt.Println(log)
	return
}
