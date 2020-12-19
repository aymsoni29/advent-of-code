package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Part1(filename string, right, down int) int {
	trees := 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	index := 0
	current := down
	for scanner.Scan() {
		line := scanner.Text()
		if current == 0 {
			index += right
			if index >= len(line) {
				index = index - len(line)
			}
			if string(line[index]) == "#" {
				trees++
			}
			current = down
		}
		current--
	}
	return trees
}

type TestCaseDay3Part2 struct {
	right int
	down  int
}

func Part2() int {
	cases := []TestCaseDay3Part2{
		{
			right: 1,
			down:  1,
		},
		{
			right: 3,
			down:  1,
		},
		{
			right: 5,
			down:  1,
		},
		{
			right: 7,
			down:  1,
		},
		{
			right: 1,
			down:  2,
		},
	}
	ans := 1
	for i := range cases {
		count := Part1("input.txt", cases[i].right, cases[i].down)
		fmt.Println(count)
		ans = ans * count
	}
	return ans
}
