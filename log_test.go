package logger

import (
	"reflect"
	"testing"
)

// ? I am yet unable to get this test working. Any ideas
func TestNewPencil(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want Pencil
	}{
		{
			"Test case: A",
			args{"./tests/mark.txt"},
			Pencil{filePath: "./tests/mark.txt",},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPencil(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPencil() = %v, want %v", got, tt.want)
			}
		})
	}
}
