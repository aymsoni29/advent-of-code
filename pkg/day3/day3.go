package day3

import (
	"bufio"
	"log"
	"os"
)

func Part1(filename string) int {
	trees := 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	skip := true
	index := 0
	for scanner.Scan() {
		line := scanner.Text()
		if !skip {
			if index >= len(line) {
				index = index - len(line)
			}
			if string(line[index]) == "#" {
				trees++
			}
		}
		index += 3
		skip = false
	}
	return trees
}
