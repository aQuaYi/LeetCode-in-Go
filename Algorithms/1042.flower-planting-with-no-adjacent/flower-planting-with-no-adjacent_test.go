package problem1042

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	N     int
	paths [][]int
	ans   []int
}{

	{
		5,
		[][]int{{4, 1}, {4, 2}, {4, 3}, {2, 5}, {1, 2}, {1, 5}},
		[]int{1, 2, 1, 3, 3},
	},

	{
		3,
		[][]int{{1, 2}, {2, 3}, {3, 1}},
		[]int{1, 2, 3},
	},

	{
		4,
		[][]int{{1, 2}, {3, 4}},
		[]int{1, 2, 1, 2},
	},

	{
		4,
		[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 3}, {2, 4}},
		[]int{1, 2, 3, 4},
	},

	// 可以有多个 testcase
}

func Test_gardenNoAdj(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, gardenNoAdj(tc.N, tc.paths), "输入:%v", tc)
	}
}

func Benchmark_gardenNoAdj(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			gardenNoAdj(tc.N, tc.paths)
		}
	}
}
