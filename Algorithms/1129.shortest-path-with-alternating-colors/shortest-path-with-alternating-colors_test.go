package problem1129

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n     int
	reds  [][]int
	blues [][]int
	ans   []int
}{

	{
		3,
		[][]int{{0, 1}, {1, 2}},
		[][]int{},
		[]int{0, 1, -1},
	},

	{
		3,
		[][]int{{0, 1}},
		[][]int{{2, 1}},
		[]int{0, 1, -1},
	},

	{
		3,
		[][]int{{1, 0}},
		[][]int{{2, 1}},
		[]int{0, -1, -1},
	},

	{
		3,
		[][]int{{0, 1}},
		[][]int{{1, 2}},
		[]int{0, 1, 2},
	},

	{
		3,
		[][]int{{0, 1}, {0, 2}},
		[][]int{{1, 0}},
		[]int{0, 1, 1},
	},

	// 可以有多个 testcase
}

func Test_shortestAlternatingPaths(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, shortestAlternatingPaths(tc.n, tc.reds, tc.blues), "输入:%v", tc)
	}
}

func Benchmark_shortestAlternatingPaths(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			shortestAlternatingPaths(tc.n, tc.reds, tc.blues)
		}
	}
}
