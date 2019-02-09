package problem0980

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	grid [][]int
	ans  int
}{

	{
		[][]int{
			{0, 1},
			{2, 0},
		},
		0,
	},

	{
		[][]int{
			{1, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 2, -1},
		},
		2,
	},

	{
		[][]int{
			{1, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 2},
		},
		4,
	},

	// 可以有多个 testcase
}

func Test_uniquePathsIII(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, uniquePathsIII(tc.grid), "输入:%v", tc)
	}
}

func Benchmark_uniquePathsIII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			uniquePathsIII(tc.grid)
		}
	}
}
