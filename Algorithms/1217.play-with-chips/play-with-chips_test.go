package problem1217

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	chips []int
	ans   int
}{

	{
		[]int{1, 2, 3},
		1,
	},

	{
		[]int{2, 2, 2, 3, 3},
		2,
	},

	// 可以有多个 testcase
}

func Test_minCostToMoveChips(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, minCostToMoveChips(tc.chips), "输入:%v", tc)
	}
}

func Benchmark_minCostToMoveChips(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minCostToMoveChips(tc.chips)
		}
	}
}
