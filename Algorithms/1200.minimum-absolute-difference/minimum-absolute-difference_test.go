package problem1200

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	arr []int
	ans [][]int
}{

	{
		[]int{4, 2, 1, 3},
		[][]int{{1, 2}, {2, 3}, {3, 4}},
	},

	{
		[]int{1, 3, 6, 10, 15},
		[][]int{{1, 3}},
	},

	{
		[]int{3, 8, -10, 23, 19, -4, -14, 27},
		[][]int{{-14, -10}, {19, 23}, {23, 27}},
	},

	// 可以有多个 testcase
}

func Test_minimumAbsDifference(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, minimumAbsDifference(tc.arr), "输入:%v", tc)
	}
}

func Benchmark_minimumAbsDifference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minimumAbsDifference(tc.arr)
		}
	}
}
