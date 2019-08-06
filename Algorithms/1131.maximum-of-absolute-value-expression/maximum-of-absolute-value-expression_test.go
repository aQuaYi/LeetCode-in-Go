package problem1131

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	B   []int
	ans int
}{

	{
		[]int{1, 2, 3, 4},
		[]int{-1, 4, 5, 6},
		13,
	},

	{
		[]int{1, -2, -5, 0, 10},
		[]int{0, -2, -1, -7, -4},
		20,
	},

	// 可以有多个 testcase
}

func Test_maxAbsValExpr(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, maxAbsValExpr(tc.A, tc.B), "输入:%v", tc)
	}
}

func Benchmark_maxAbsValExpr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxAbsValExpr(tc.A, tc.B)
		}
	}
}
