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

package main

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
			"A",
			args{"Test A"},
			fmt.Sprintf("%v: %v", time.Now().Format(textF), "Test A"),
		},
		{
			"B",
			args{"Test B"},
			fmt.Sprintf("%v: %v", time.Now().Format(textF), "Test B"),
		},
		{
			"C",
			args{"Test C"},
			fmt.Sprintf("%v: %v", time.Now().Format(textF), "Test C"),
		},
		{
			"D",
			args{"Test D"},
			fmt.Sprintf("%v: %v", time.Now().Format(textF), "Test D"),
		},
		{
			"E",
			args{"Test E"},
			fmt.Sprintf("%v: %v", time.Now().Format(textF), "Test E"),
		},
		{
			"F",
			args{"Test F"},
			fmt.Sprintf("%v: %v", time.Now().Format(textF), "Test F"),
		},
		{
			"G",
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
