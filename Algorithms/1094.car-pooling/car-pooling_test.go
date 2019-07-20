package problem1094

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	trips    [][]int
	capacity int
	ans      bool
}{

	{
		[][]int{{2, 1, 5}, {3, 3, 7}},
		4,
		false,
	},

	{
		[][]int{{2, 1, 5}, {3, 3, 7}},
		5,
		true,
	},

	{
		[][]int{{2, 1, 5}, {3, 5, 7}},
		3,
		true,
	},

	{
		[][]int{{3, 2, 7}, {3, 7, 9}, {8, 3, 9}},
		11,
		true,
	},

	// 可以有多个 testcase
}

func Test_carPooling(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, carPooling(tc.trips, tc.capacity), "输入:%v", tc)
	}
}

func Benchmark_carPooling(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			carPooling(tc.trips, tc.capacity)
		}
	}
}
