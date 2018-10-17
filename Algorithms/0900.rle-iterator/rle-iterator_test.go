package problem0900

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RLEIterator(t *testing.T) {
	ast := assert.New(t)

	ints := []int{3, 8, 0, 9, 2, 5}

	r := Constructor(ints)

	nums := []int{2, 1, 1, 2}
	expecteds := []int{8, 8, 5, -1}

	for i, n := range nums {
		expected := expecteds[i]
		actual := r.Next(n)
		ast.Equal(expected, actual)
	}
}
