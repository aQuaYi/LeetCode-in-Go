package problem0703

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_kthLargest(t *testing.T) {
	ast := assert.New(t)
	k := 3
	nums := []int{4, 5, 8, 2}

	kl := Constructor(k, nums)

	adds := []int{3, 5, 10, 9, 4}
	rets := []int{4, 5, 5, 8, 8}

	for i := range adds {
		expected := rets[i]
		actual := kl.Add(adds[i])
		ast.Equal(expected, actual, "kl.Add(%d) != %d\n", adds[i], expected)
	}
}

func Test_kthLargest_2(t *testing.T) {
	ast := assert.New(t)
	k := 1
	nums := []int{}

	kl := Constructor(k, nums)

	adds := []int{-3, -2, -4, 0, 4}
	rets := []int{-3, -2, -2, 0, 4}

	for i := range adds {
		expected := rets[i]
		actual := kl.Add(adds[i])
		ast.Equal(expected, actual, "kl.Add(%d) != %d\n", adds[i], expected)
	}
}
