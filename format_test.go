package logger

import (
	"fmt"
	"testing"
	"time"
)

func Test_format(t *testing.T) {
	var textF string = "2006-01-02 15:04:05"
	type args struct {
		log string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Test case: A",
			args{"Test A"},
			fmt.Sprintf("%v: %v", time.Now().Format(textF), "Test A"),
		},
		{
			"Test case: B",
			args{"Test B"},
			fmt.Sprintf("%v: %v", time.Now().Format(textF), "Test B"),
		},
		{
			"Test case: C",
			args{"Test C"},
			fmt.Sprintf("%v: %v", time.Now().Format(textF), "Test C"),
		},
		{
			"Test case: D",
			args{"Test D"},
			fmt.Sprintf("%v: %v", time.Now().Format(textF), "Test D"),
		},
		{
			"Test case: E",
			args{"Test E"},
			fmt.Sprintf("%v: %v", time.Now().Format(textF), "Test E"),
		},
		{
			"Test case: F",
			args{"Test F"},
			fmt.Sprintf("%v: %v", time.Now().Format(textF), "Test F"),
		},
		{
			"Test case: G",
			args{"Test G"},
			fmt.Sprintf("%v: %v", time.Now().Format(textF), "Test G"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := format(tt.args.log); got != tt.want {
				t.Errorf("format() = %v, want %v", got, tt.want)
			}
		})
	}
}
