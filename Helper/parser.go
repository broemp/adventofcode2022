package Helper

import (
	"bufio"
	"log"
	"os"
)

func ParseFile(path string) []string {
	// Open file
	file, err := os.Open(path)
	// Check for errors
	if err != nil {
		log.Fatal(err)
	}
	// Close file when done
	defer file.Close()
	// Create scanner
	scanner := bufio.NewScanner(file)
	// Create slice to store lines
	var lines []string
	// Scan lines
	for scanner.Scan() {
		// Append line to slice
		lines = append(lines, scanner.Text())
	}
	// Check for errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// Return slice
	return lines
}
