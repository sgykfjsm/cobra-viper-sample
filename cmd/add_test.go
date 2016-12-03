package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	var addNormalTests = []struct {
		nums     []int
		x        int
		expected int
	}{
		{[]int{}, 0, 0},
		{[]int{1}, 0, 1},
		{[]int{1, 2}, 0, 3},
	}

	var actual int
	for _, an := range addNormalTests {
		actual = add(an.nums, an.x)
		assert.Equal(t, actual, an.expected)
	}
}
