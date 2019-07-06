package problem1049

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	stones []int
	ans    int
}{

	{
		[]int{3},
		3,
	},

	{
		[]int{1, 2, 3},
		0,
	},

	{
		[]int{31, 26, 33, 21, 40},
		5,
	},

	{
		[]int{21, 16, 23, 32, 25, 13, 20, 18, 22, 21, 84, 35, 33, 17, 27, 24, 10, 19, 31, 26, 94, 37, 31, 25, 24, 25, 15, 23, 17, 13},
		1,
	},

	{
		[]int{3, 2, 1, 7, 3, 7},
		1,
	},

	{
		[]int{2, 7, 4, 1, 8, 1},
		1,
	},

	// 可以有多个 testcase
}

func Test_lastStoneWeightII(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, lastStoneWeightII(tc.stones), "输入:%v", tc)
	}
}

func Benchmark_lastStoneWeightII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			lastStoneWeightII(tc.stones)
		}
	}
}
