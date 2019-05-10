package problem0994

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
		[][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}},
		4,
	},

	{
		[][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}},
		-1,
	},

	{
		[][]int{{0, 2}},
		0,
	},

	{
		[][]int{{0}},
		0,
	},

	// 可以有多个 testcase
}

func Test_orangesRotting(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, orangesRotting(tc.grid), "输入:%v", tc)
	}
}

func Benchmark_orangesRotting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			orangesRotting(tc.grid)
		}
	}
}
