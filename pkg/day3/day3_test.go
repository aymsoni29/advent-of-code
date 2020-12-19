package day3

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_Part1(t *testing.T) {
	assert.Equal(t, 178, Part1("input.txt", 3, 1))
}

func Test_Part2(t *testing.T) {
	assert.Equal(t, 3492520200, Part2())
}
