package problem0394

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	ans string
}{

	{
		"aaabcbc",
		"aaabcbc",
	},

	{
		"10[a]2[bc]",
		"aaaaaaaaaabcbc",
	},

	{
		"3[a]2[bc]",
		"aaabcbc",
	},

	{
		"3[a2[c]]",
		"accaccacc",
	},

	{
		"2[abc]3[cd]ef",
		"abcabccdcdcdef",
	},

	// 可以有多个 testcase
}

func Test_decodeString(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, decodeString(tc.s), "输入:%v", tc)
	}
}

func Benchmark_decodeString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			decodeString(tc.s)
		}
	}
}
