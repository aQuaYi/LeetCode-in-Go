package problem1124

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	hours []int
	ans   int
}{

	{
		[]int{16, 15, 11, 2, 9, 3},
		6,
	},

	{
		[]int{9, 9, 6, 0, 6, 6, 9},
		3,
	},

	{
		[]int{6, 6, 6, 6, 9, 9, 9, 9, 9, 9, 6, 0, 6, 6, 9},
		13,
	},

	// 可以有多个 testcase
}

func Test_longestWPI(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, longestWPI(tc.hours), "输入:%v", tc)
	}
}

func Benchmark_longestWPI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			longestWPI(tc.hours)
		}
	}
}
