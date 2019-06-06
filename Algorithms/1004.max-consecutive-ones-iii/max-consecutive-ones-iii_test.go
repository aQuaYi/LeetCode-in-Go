package problem1004

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	K   int
	ans int
}{

	{
		[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
		2,
		6,
	},

	{
		[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
		3,
		10,
	},

	// 可以有多个 testcase
}

func Test_longestOnes(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, longestOnes(tc.A, tc.K), "输入:%v", tc)
	}
}

func Benchmark_longestOnes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			longestOnes(tc.A, tc.K)
		}
	}
}
