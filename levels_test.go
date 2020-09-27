package logger

import (
	"fmt"
	"time"
	"testing"
)

// func main() {
// 	Log("This is a dumb thing")
// }

func TestLog(t *testing.T) {
	type args struct {
		i string
	}
	tests := []struct {
		name    string
		args    args
		wantLog string
	}{
		{
			"Test case: A",
			args{"Test A"},
			fmt.Sprintf("%v: %v", time.Now().Format("2006-01-02 15:04:05"), "Test A"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLog := Log(tt.args.i); gotLog != tt.wantLog {
				t.Errorf("Log() = %v, want %v", gotLog, tt.wantLog)
			}
		})
	}
}
