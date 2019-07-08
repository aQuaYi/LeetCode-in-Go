package problem1053

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
		[]int{1, 1, 5},
		[]int{1, 1, 5},
	},

	{
		[]int{2, 1, 1},
		[]int{1, 2, 1},
	},

	{
		[]int{3, 2, 1},
		[]int{3, 1, 2},
	},

	{
		[]int{1, 9, 4, 6, 7},
		[]int{1, 7, 4, 6, 9},
	},

	{
		[]int{3, 1, 1, 3},
		[]int{1, 3, 1, 3},
	},

	// 可以有多个 testcase
}

func Test_prevPermOpt1(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, prevPermOpt1(tc.A), "输入:%v", tc)
	}
}

func Benchmark_prevPermOpt1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			prevPermOpt1(tc.A)
		}
	}
}
