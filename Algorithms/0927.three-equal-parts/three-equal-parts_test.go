package problem0927

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	ans []int
}{

	{
		[]int{0, 0, 0, 0, 0},
		[]int{0, 2},
	},

	{
		[]int{1, 0, 1, 0, 1, 0, 0},
		[]int{-1, -1},
	},

	{
		[]int{1, 0, 1, 0, 1},
		[]int{0, 3},
	},

	{
		[]int{1, 1, 0, 1, 1},
		[]int{-1, -1},
	},

	// 可以有多个 testcase
}

func Test_threeEqualParts(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, threeEqualParts(tc.A), "输入:%v", tc)
	}
}

func Benchmark_threeEqualParts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			threeEqualParts(tc.A)
		}
	}
}
