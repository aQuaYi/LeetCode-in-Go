package problem1130

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
		[]int{12, 4, 13, 8, 10, 1, 5, 14, 5, 4, 3, 11, 6, 7, 5, 12, 6, 4, 14, 15, 2, 2, 11, 13, 10, 15, 3, 11, 15, 1, 2, 10, 6, 14, 15, 9, 7, 9, 9, 5},
		3898,
	},

	{
		[]int{1, 2, 3, 4, 5},
		40,
	},

	{
		[]int{6, 2, 4},
		32,
	},

	// 可以有多个 testcase
}

func Test_mctFromLeafValues(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, mctFromLeafValues(tc.arr), "输入:%v", tc)
	}
}

func Benchmark_mctFromLeafValues(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			mctFromLeafValues(tc.arr)
		}
	}
}
