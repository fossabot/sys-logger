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
	"os"
	"errors"
)

// Pencil data type will be used
// for editing lines from the log
// ! Don't use this type directly.
// ! Instead use NewPencil().
type Pencil struct {
	filePath     string // File path
	addition     uint32 // Total lines added
	subtractions uint32 // Total lines removed
}

// * Getters and setters below.

// NewPencil returns a new Pencil or an error.
func NewPencil(path string) (Pencil, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return Pencil{}, errors.New("Cannot read file. Please provide a file with read and write privileges")
	}
	return Pencil{
		filePath: path,
	}, nil
}
