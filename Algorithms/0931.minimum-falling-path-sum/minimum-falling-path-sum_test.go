package problem0931

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   [][]int
	ans int
}{

	{
		[][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		12,
	},

	// 可以有多个 testcase
}

func Test_minFallingPathSum(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, minFallingPathSum(tc.A), "输入:%v", tc)
	}
}

func Benchmark_minFallingPathSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minFallingPathSum(tc.A)
		}
	}
}
