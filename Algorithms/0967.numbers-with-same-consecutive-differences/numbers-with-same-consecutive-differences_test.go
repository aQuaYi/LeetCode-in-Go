package problem0967

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	N   int
	K   int
	ans []int
}{

	{
		3,
		1,
		[]int{101, 121, 123, 210, 212, 232, 234, 321, 323, 343, 345, 432, 434, 454, 456, 543, 545, 565, 567, 654, 656, 676, 678, 765, 767, 787, 789, 876, 878, 898, 987, 989},
	},

	{
		2,
		0,
		[]int{11, 22, 33, 44, 55, 66, 77, 88, 99},
	},

	{
		1,
		6,
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	},

	{
		1,
		0,
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	},

	{
		3,
		7,
		[]int{181, 292, 707, 818, 929},
	},

	{
		2,
		1,
		[]int{10, 12, 21, 23, 32, 34, 43, 45, 54, 56, 65, 67, 76, 78, 87, 89, 98},
	},

	// 可以有多个 testcase
}

func Test_numsSameConsecDiff(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, numsSameConsecDiff(tc.N, tc.K), "输入:%v", tc)
	}
}

func Benchmark_numsSameConsecDiff(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numsSameConsecDiff(tc.N, tc.K)
		}
	}
}
