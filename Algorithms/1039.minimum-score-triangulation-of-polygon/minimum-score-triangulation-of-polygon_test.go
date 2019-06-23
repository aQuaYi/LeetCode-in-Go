package problem1039

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	ans int
}{

	{
		[]int{1, 2, 3, 4, 1, 6, 7, 8, 1, 10, 12},
		255,
	},

	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		238,
	},

	{
		[]int{2, 2, 2, 2, 1},
		12,
	},

	{
		[]int{1, 3, 1, 4, 1, 5},
		13,
	},

	{
		[]int{1, 2, 3},
		6,
	},

	{
		[]int{3, 7, 4, 5},
		144,
	},

	// 可以有多个 testcase
}

func Test_minScoreTriangulation(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, minScoreTriangulation(tc.A), "输入:%v", tc)
	}
}

func Benchmark_minScoreTriangulation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minScoreTriangulation(tc.A)
		}
	}
}
