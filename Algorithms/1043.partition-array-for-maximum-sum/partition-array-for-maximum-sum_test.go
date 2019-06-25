package problem1043

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	K   int
	ans int
}{

	{
		[]int{10, 9, 3, 2},
		2,
		30,
	},

	{
		[]int{1, 15, 7, 9, 2, 5, 10},
		3,
		84,
	},

	// 可以有多个 testcase
}

func Test_maxSumAfterPartitioning(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, maxSumAfterPartitioning(tc.A, tc.K), "输入:%v", tc)
	}
}

func Benchmark_maxSumAfterPartitioning(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxSumAfterPartitioning(tc.A, tc.K)
		}
	}
}
