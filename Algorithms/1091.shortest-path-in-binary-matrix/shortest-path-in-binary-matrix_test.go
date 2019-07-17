package problem1091

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
		[][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}},
		-1,
	},

	{
		[][]int{{0, 1}, {1, 0}},
		2,
	},

	{
		[][]int{{0, 1, 0}, {1, 1, 0}, {1, 1, 0}},
		-1,
	},

	{
		[][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}},
		4,
	},

	// 可以有多个 testcase
}

func Test_shortestPathBinaryMatrix(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, shortestPathBinaryMatrix(tc.grid), "输入:%v", tc)
	}
}

func Benchmark_shortestPathBinaryMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			shortestPathBinaryMatrix(tc.grid)
		}
	}
}
