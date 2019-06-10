package problem1014

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
		[]int{1, 2},
		2,
	},

	{
		[]int{8, 1, 5, 2, 6},
		11,
	},

	// 可以有多个 testcase
}

func Test_maxScoreSightseeingPair(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, maxScoreSightseeingPair(tc.A), "输入:%v", tc)
	}
}

func Benchmark_maxScoreSightseeingPair(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxScoreSightseeingPair(tc.A)
		}
	}
}
