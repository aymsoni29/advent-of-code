package day4

import (
	"bufio"
	"log"
	"os"
)

func findValid(filename string) {

	validCount := 0

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {

		} else {

		}
	}
}

func Split(r rune) bool {
	return r == ':' || r == ' '
}
