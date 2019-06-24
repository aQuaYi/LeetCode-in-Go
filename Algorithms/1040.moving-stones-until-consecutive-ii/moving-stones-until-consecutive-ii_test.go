package problem1040

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	stones []int
	ans    []int
}{

	{
		[]int{6, 5, 4, 3, 10},
		[]int{2, 3},
	},

	{
		[]int{1, 2, 3, 4, 5, 111, 112, 113, 114, 115},
		[]int{5, 105},
	},

	{
		[]int{1, 2, 3, 4, 5, 11, 12, 13, 14, 15},
		[]int{5, 5},
	},

	{
		[]int{1, 3, 900, 904, 1000},
		[]int{3, 994},
	},

	{
		[]int{1, 50, 900, 905, 170000},
		[]int{4, 169947},
	},

	{
		[]int{1, 50, 900, 904, 170000},
		[]int{3, 169947},
	},

	{
		[]int{1, 50, 900, 13000, 170000},
		[]int{4, 169947},
	},

	{
		[]int{4, 8, 9},
		[]int{2, 3},
	},

	{
		[]int{1, 5, 9},
		[]int{2, 3},
	},

	{
		[]int{1, 5, 9, 13},
		[]int{3, 6},
	},

	{
		[]int{1, 50, 900, 13000},
		[]int{3, 12948},
	},

	{
		[]int{7, 4, 9},
		[]int{1, 2},
	},

	{
		[]int{100, 101, 104, 102, 103},
		[]int{0, 0},
	},

	// 可以有多个 testcase
}

func Test_numMovesStonesII(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, numMovesStonesII(tc.stones), "输入:%v", tc)
	}
}

func Benchmark_numMovesStonesII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numMovesStonesII(tc.stones)
		}
	}
}
