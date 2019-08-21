package problem1162

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
		[][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
		-1,
	},

	{
		[][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		2,
	},

	{
		[][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		4,
	},

	// 可以有多个 testcase
}

func Test_maxDistance(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, maxDistance(tc.grid), "输入:%v", tc)
	}
}

func Benchmark_maxDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxDistance(tc.grid)
		}
	}
}
