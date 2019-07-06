package problem1044

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	S   string
	ans string
}{

	{
		"banana",
		"ana",
	},

	{
		"abcd",
		"",
	},

	// 可以有多个 testcase
}

func Test_longestDupSubstring(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, longestDupSubstring(tc.S), "输入:%v", tc)
	}
}

func Benchmark_longestDupSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			longestDupSubstring(tc.S)
		}
	}
}
