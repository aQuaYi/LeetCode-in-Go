package problem1027

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	ans int
}{

	{
		[]int{44, 46, 22, 68, 45, 66, 43, 9, 37, 30, 50, 67, 32, 47, 44, 11, 15, 4, 11, 6, 20, 64, 54, 54, 61, 63, 23, 43, 3, 12, 51, 61, 16, 57, 14, 12, 55, 17, 18, 25, 19, 28, 45, 56, 29, 39, 52, 8, 1, 21, 17, 21, 23, 70, 51, 61, 21, 52, 25, 28},
		6,
	},

	{
		[]int{3, 6, 9, 12},
		4,
	},

	{
		[]int{9, 4, 7, 2, 10},
		3,
	},

	{
		[]int{20, 1, 15, 3, 10, 5, 8},
		4,
	},

	// 可以有多个 testcase
}

func Test_longestArithSeqLength(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, longestArithSeqLength(tc.A), "输入:%v", tc)
	}
}

func Benchmark_longestArithSeqLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			longestArithSeqLength(tc.A)
		}
	}
}
