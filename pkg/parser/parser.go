package parser

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// ParseTextFileToInt takes in a file and returns an array of integer
func ParseTextFileToInt(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var output []int

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("Error parsing input")
		}
		output = append(output, num)
	}

	return output
}
