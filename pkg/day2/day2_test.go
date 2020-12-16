package day2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Part1(t *testing.T) {
	assert.Equal(t, 467, Part1("../day2/input.txt"))
}

func Test_Part2(t *testing.T) {
	assert.Equal(t, 441, Part2("../day2/input.txt"))
}
