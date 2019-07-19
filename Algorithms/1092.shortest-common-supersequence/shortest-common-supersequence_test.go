package problem1092

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	str1 string
	str2 string
	ans  string
}{

	{
		"abaa",
		"aaaa",
		"aabaa",
	},

	{
		"cab",
		"abac",
		"cabac",
	},

	{
		"abac",
		"cab",
		"cabac",
	},

	// 可以有多个 testcase
}

func Test_shortestCommonSupersequence(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, shortestCommonSupersequence(tc.str1, tc.str2), "输入:%v", tc)
	}
}

func Benchmark_shortestCommonSupersequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			shortestCommonSupersequence(tc.str1, tc.str2)
		}
	}
}
