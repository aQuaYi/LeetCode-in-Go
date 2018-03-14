package problem0214

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
		"abaa",
		"aabaa",
	},

	{
		"a",
		"a",
	},

	{
		"abcba",
		"abcba",
	},

	{
		"aacecaaa",
		"aaacecaaa",
	},

	{
		"abcd",
		"dcbabcd",
	},

	// 可以有多个 testcase
}

func Test_shortestPalindrome(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, shortestPalindrome(tc.s), "输入:%v", tc)
	}
}

func Benchmark_shortestPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			shortestPalindrome(tc.s)
		}
	}
}
