/*
Copyright 2020 Abhigan Kumar

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

https://www.eclipse.org/legal/epl-2.0/

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
or implied. See the License for the specific language governing
permissions and limitations under the License.
*/

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
			"Alpha",
			args{"./tests/log/alpha.txt"},
			Pencil{filePath: "./tests/log/alpha.txt",},
		},
		{
			"Beta",
			args{"./tests/log/beta.txt"},
			Pencil{filePath: "./tests/log/beta.txt",},
		},
		{
			"Gamma",
			args{"./tests/log/gamma.txt"},
			Pencil{filePath: "./tests/log/gamma.txt",},
		},
		{
			"Lambda",
			args{"./tests/log/lambda.txt"},
			Pencil{filePath: "./tests/log/lambda.txt",},
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
