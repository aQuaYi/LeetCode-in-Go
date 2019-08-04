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
		[][]int{{1, 1}},
		[]int{0, 1, 3},
	},

	{
		9,
		[][]int{{3, 1}, {2, 3}, {7, 6}, {5, 1}, {1, 3}, {8, 1}, {5, 4}, {8, 4}, {6, 3}, {4, 7}, {0, 1}, {7, 8}, {3, 8}},
		[][]int{{4, 1}, {5, 8}, {3, 7}, {7, 1}, {1, 8}, {8, 7}, {5, 4}},
		[]int{0, 1, -1, 5, 3, -1, 7, 6, 2},
	},

	{
		5,
		[][]int{{0, 1}, {1, 3}, {2, 1}, {3, 1}, {2, 0}, {1, 0}, {0, 2}},
		[][]int{{1, 2}, {3, 2}, {2, 4}, {3, 3}, {0, 3}, {1, 4}, {0, 1}, {0, 2}, {0, 0}, {4, 1}},
		[]int{0, 1, 1, 1, 2},
	},

	{
		5,
		[][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
		[][]int{{1, 2}, {2, 3}, {3, 1}},
		[]int{0, 1, 2, 3, 7},
	},

	{
		3,
		[][]int{{1, 2}},
		[][]int{{0, 1}},
		[]int{0, 1, 2},
	},

	{
		3,
		[][]int{{0, 1}},
		[][]int{{1, 2}},
		[]int{0, 1, 2},
	},

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
