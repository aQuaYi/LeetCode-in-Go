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
