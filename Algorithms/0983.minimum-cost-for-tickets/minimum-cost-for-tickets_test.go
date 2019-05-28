package problem0983

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	days  []int
	costs []int
	ans   int
}{

	{
		[]int{1, 4, 6, 7, 8, 20},
		[]int{2, 7, 15},
		11,
	},

	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31},
		[]int{2, 7, 15},
		17,
	},

	// 可以有多个 testcase
}

func Test_mincostTickets(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, mincostTickets(tc.days, tc.costs), "输入:%v", tc)
	}
}

func Benchmark_mincostTickets(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			mincostTickets(tc.days, tc.costs)
		}
	}
}
