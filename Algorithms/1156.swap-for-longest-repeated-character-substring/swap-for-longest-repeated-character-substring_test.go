package problem1156

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	text string
	ans  int
}{

	{
		"ababa",
		3,
	},

	{
		"aaabaaa",
		6,
	},

	{
		"aaabbbbbaaa",
		5,
	},

	{
		"aaabbaaa",
		4,
	},

	{
		"aaaaa",
		5,
	},

	{
		"abcdef",
		1,
	},

	// 可以有多个 testcase
}

func Test_maxRepOpt1(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, maxRepOpt1(tc.text), "输入:%v", tc)
	}
}

func Benchmark_maxRepOpt1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxRepOpt1(tc.text)
		}
	}
}
