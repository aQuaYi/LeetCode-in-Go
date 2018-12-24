package problem0940

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	S   string
	ans int
}{

	{
		"adsflkasldgalksdglakdsfjlakjdsglkasjdlkasdadsflkdsfldaslfjsdaldadsabc",
		589422165,
	},

	{
		"abc",
		7,
	},

	{
		"aba",
		6,
	},

	{
		"aaa",
		3,
	},

	// 可以有多个 testcase
}

func Test_distinctSubseqII(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, distinctSubseqII(tc.S), "输入:%v", tc)
	}
}

func Benchmark_distinctSubseqII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			distinctSubseqII(tc.S)
		}
	}
}
