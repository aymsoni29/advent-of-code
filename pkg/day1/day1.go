package day1

import (
	"fmt"
	"sort"
)

// Part1 of Day 1
// Approach 1 - sort + two pointers = O(nlog(n))
// Approach 2 - HashMap (value -> occurences) O(n)
// Approach 3 - Two for loops (every num[i] iterate through all entries) O(n^2)
// Approach 4 - For loop + Binary Search (O(nlog(n)))
func Part1(input []int) int {
	fmt.Println("Day 1 : Part 1")

	approach := 1
	switch approach {
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
				return i*j
			} else if input[i]+input[j] < 2020 {
				i++
			} else {
				j--
			}
		}
	case 2:
		fmt.Println("Approach: ", approach)
	case 3:
		fmt.Println("Approach: ", approach)
	case 4:
		fmt.Println("Approach: ", approach)
	}

	return -1
}
