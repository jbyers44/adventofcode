package utils

import (
	"log"
	"os"
)

func OpenFile(filepath string) *os.File {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}

	return file
}
