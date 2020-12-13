package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Part1(t *testing.T) {
	assert.Equal(t, 326211, Part1("../day1/input.txt"))
}
