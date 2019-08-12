package problem1143

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   string
	B   string
	ans int
}{

	{
		"abcde",
		"ace",
		3,
	},

	{
		"abc",
		"abc",
		3,
	},

	{
		"abc",
		"def",
		0,
	},

	// 可以有多个 testcase
}

func Test_longestCommonSubsequence(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, longestCommonSubsequence(tc.A, tc.B), "输入:%v", tc)
	}
}

func Benchmark_longestCommonSubsequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			longestCommonSubsequence(tc.A, tc.B)
		}
	}
}
