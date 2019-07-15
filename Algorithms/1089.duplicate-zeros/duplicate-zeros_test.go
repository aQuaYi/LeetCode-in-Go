package problem1089

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	arr []int
	ans []int
}{

	{

		[]int{8, 4, 5, 0, 0, 0, 0, 7},
		[]int{8, 4, 5, 0, 0, 0, 0, 0},
	},

	{
		[]int{0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0},
	},

	{
		[]int{0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0},
	},

	{
		[]int{1, 0, 2, 3, 0, 4, 5, 0},
		[]int{1, 0, 0, 2, 3, 0, 0, 4},
	},

	{
		[]int{1, 2, 3},
		[]int{1, 2, 3},
	},

	// 可以有多个 testcase
}

func Test_duplicateZeros(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		duplicateZeros(tc.arr)
		ast.Equal(tc.ans, tc.arr, "输入:%v", tc)
	}
}

func Benchmark_duplicateZeros(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			duplicateZeros(tc.arr)
		}
	}
}
