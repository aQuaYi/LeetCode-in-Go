package problem0977

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	ans []int
}{

	{
		[]int{-4, -1, 0, 3, 10},
		[]int{0, 1, 9, 16, 100},
	},

	{
		[]int{-7, -3, 2, 3, 11},
		[]int{4, 9, 9, 49, 121},
	},

	// 可以有多个 testcase
}

func Test_sortedSquares(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, sortedSquares(tc.A), "输入:%v", tc)
	}
}

func Benchmark_sortedSquares(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			sortedSquares(tc.A)
		}
	}
}
