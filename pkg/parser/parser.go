package parser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// StructDay2 holds the necessary data for each test case
type StructDay2 struct {
	Min  int
	Max  int
	Char string
	Pass string
}

// ParseTextFileToInt takes in a file and returns an array of integer
func ParseTextFileToInt(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
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

func ParseDay2Input(filename string) []StructDay2 {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Error reading file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var output []StructDay2

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.FieldsFunc(line, Split)

		min, err := strconv.Atoi(fields[0])
		if err != nil {
			fmt.Println("Failed to parse string input")
		}

		max, err := strconv.Atoi(fields[1])
		if err != nil {
			fmt.Println("Failed to parse string input")
		}

		character := fields[2]
		password := fields[3]

		data := StructDay2{
			Min:  min,
			Max:  max,
			Char: character,
			Pass: password,
		}

		output = append(output, data)
	}

	return output
}

func Split(r rune) bool {
	return r == '-' || r == ':' || r == ' '
}
