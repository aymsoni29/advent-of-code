package day1

import (
	"fmt"
	"sort"

	"github.com/advent-of-code/pkg/parser"
)

// Part1 of Day 1
// Approach 1 - sort + two pointers = O(nlog(n))
// Approach 2 - HashMap (value -> occurences) O(n)
// Approach 3 - Two for loops (every num[i] iterate through all entries) O(n^2)
// Approach 4 - For loop + Binary Search (O(nlog(n)))
func Part1(filename string) int {
	fmt.Println("Day 1 : Part 1")

	input := parser.ParseTextFileToInt(filename)
	switch approach := 2; approach {
	case 1:
		fmt.Println("Approach: ", approach)

		// Sort the input first - O(nlog(n))
		sort.Slice(input, func(i, j int) bool {
			return input[i] < input[j]
		})

		i := 0
		j := len(input) - 1
		for i < j {
			if input[i]+input[j] == 2020 {
				return input[i] * input[j]
			} else if input[i]+input[j] < 2020 {
				i++
			} else {
				j--
			}
		}
	case 2:
		fmt.Println("Approach: ", approach)

		inputByCount := make(map[int]int)
		for _, val := range input {
			inputByCount[val]++
		}

		for _, val := range input {
			if 2020-val == val && inputByCount[val] >= 2 {
				return val * val
			} else if count, ok := inputByCount[2020-val]; ok && count >= 1 {
				fmt.Println(val)
				return (2020 - val) * val
			}
		}
		return -1
	case 3:
		fmt.Println("Approach: ", approach)
		// TODO
	case 4:
		fmt.Println("Approach: ", approach)
		// TODO
	}

	return -1
}

// Part2 of Day1
func Part2(filename string) int {
	fmt.Println("Day 1 : Part 2")

	input := parser.ParseTextFileToInt(filename)

	inputByCount := make(map[int]int)
	for _, val := range input {
		inputByCount[val]++
	}

	for i := 0; i < len(input); i++ {
		inputByCount[input[i]]--
		for j := 1; j < len(input); j++ {
			// We don't want to recount what we have in current iteration
			inputByCount[input[j]]--
			if inputByCount[2020-(input[i]+input[j])] >= 1 {
				return input[i] * input[j] * (2020 - input[i] - input[j])
			}
			// Since we are done checking
			inputByCount[input[j]]++
		}
		inputByCount[input[i]]++
	}
	return -1
}
