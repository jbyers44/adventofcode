package utils

import (
	"bufio"
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

// Read an entire file into memory, formatted as an array of strings. Only for small files.
func FileLines(file *os.File) (lines []string) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return
}
