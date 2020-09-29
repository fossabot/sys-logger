package logger

import (
	"os"
	"log"
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
func NewPencil(path string) Pencil {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatalln("Cannot read file. Please provide a file with read and write privileges.")
	}
	return Pencil{
		filePath: path,
	}
}
