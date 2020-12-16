package day2

import (
	"github.com/advent-of-code/pkg/parser"
)

// Part 1 of day 2
func Part1(filename string) int {
	input := parser.ParseDay2Input(filename)
	ans := 0
	for _, item := range input {
		letterCount := make(map[string]int)
		for _, letter := range item.Pass {
			letterCount[string(letter)]++
		}
		if letterCount[item.Char] <= item.Max && letterCount[item.Char] >= item.Min {
			ans++
		}
	}
	return ans
}
