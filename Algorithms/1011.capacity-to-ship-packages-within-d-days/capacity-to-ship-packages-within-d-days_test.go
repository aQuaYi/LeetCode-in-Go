package problem1011

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	weights []int
	D       int
	ans     int
}{

	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		5,
		15,
	},

	{
		[]int{3, 2, 2, 4, 1, 4},
		3,
		6,
	},

	{
		[]int{1, 2, 3, 1, 1},
		4,
		3,
	},

	// 可以有多个 testcase
}

func Test_shipWithinDays(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, shipWithinDays(tc.weights, tc.D), "输入:%v", tc)
	}
}

func Benchmark_shipWithinDays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			shipWithinDays(tc.weights, tc.D)
		}
	}
}
