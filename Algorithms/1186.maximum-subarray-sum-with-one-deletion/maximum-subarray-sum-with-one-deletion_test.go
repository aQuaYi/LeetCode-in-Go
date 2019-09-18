package problem1186

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	arr []int
	ans int
}{

{
	[]int{ -50 },
-50,
},

	{
[]int{ 1,-2,-2,3 },
3,
	},

	{
		[]int{11, -10, -11, 8, 7, -6, 9, 4, 11, 6, 5, 0},
		50,
	},

	{
		[]int{1, -4, -5, -2, 5, 0, -1, 2},
		7,
	},

	{
		[]int{8, -1, 6, -7, -4, 5, -4, 7, -6},
		17,
	},

	{
		[]int{1, -2, 0, 3},
		4,
	},

	{
		[]int{1, -2, -2, 3},
		3,
	},

	{
		[]int{-1, -1, -1, -1},
		-1,
	},

	// 可以有多个 testcase
}

func Test_maximumSum(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, maximumSum(tc.arr), "输入:%v", tc)
	}
}

func Benchmark_maximumSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maximumSum(tc.arr)
		}
	}
}
