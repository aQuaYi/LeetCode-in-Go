package problem1139

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
		[][]int{{1, 1, 1, 1}, {1, 0, 0, 1}, {1, 1, 1, 1}},
		1,
	},

	{
		[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
		9,
	},

	{
		[][]int{{0}},
		0,
	},

	{
		[][]int{{1, 1, 0, 0}},
		1,
	},

	// 可以有多个 testcase
}

func Test_largest1BorderedSquare(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, largest1BorderedSquare(tc.grid), "输入:%v", tc)
	}
}

func Benchmark_largest1BorderedSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			largest1BorderedSquare(tc.grid)
		}
	}
}
