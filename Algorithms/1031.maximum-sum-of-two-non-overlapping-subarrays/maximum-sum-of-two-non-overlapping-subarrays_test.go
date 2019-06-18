package problem1031

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	L   int
	M   int
	ans int
}{

	{
		[]int{0, 6, 5, 2, 2, 5, 1, 9, 4},
		1,
		2,
		20,
	},

	{
		[]int{3, 8, 1, 3, 2, 1, 8, 9, 0},
		3,
		2,
		29,
	},

	{
		[]int{2, 1, 5, 6, 0, 9, 5, 0, 3, 8},
		4,
		3,
		31,
	},

	// 可以有多个 testcase
}

func Test_maxSumTwoNoOverlap(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, maxSumTwoNoOverlap(tc.A, tc.L, tc.M), "输入:%v", tc)
	}
}

func Benchmark_maxSumTwoNoOverlap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxSumTwoNoOverlap(tc.A, tc.L, tc.M)
		}
	}
}
