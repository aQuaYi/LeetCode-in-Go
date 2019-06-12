package problem1019

import (
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	head []int
	ans  []int
}{

	{
		[]int{2, 1, 5},
		[]int{5, 5, 0},
	},

	{
		[]int{2, 7, 4, 3, 5},
		[]int{7, 0, 5, 5, 0},
	},

	{
		[]int{1, 7, 5, 1, 9, 2, 5, 1},
		[]int{7, 9, 9, 9, 0, 5, 0, 0},
	},

	// 可以有多个 testcase
}

func Test_nextLargerNodes(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		head := kit.Ints2List(tc.head)
		ast.Equal(tc.ans, nextLargerNodes(head), "输入:%v", tc)
	}
}

func Benchmark_nextLargerNodes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			head := kit.Ints2List(tc.head)
			nextLargerNodes(head)
		}
	}
}
