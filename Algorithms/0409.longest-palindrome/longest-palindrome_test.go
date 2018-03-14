package problem0409

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	ans int
}{

	{
		"zzzzzz",
		6,
	},

	{
		"AAAAAA",
		6,
	},

	{
		"abccccdd",
		7,
	},

	// 可以有多个 testcase
}

func Test_longestPalindrome(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, longestPalindrome(tc.s), "输入:%v", tc)
	}
}

func Benchmark_longestPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			longestPalindrome(tc.s)
		}
	}
}
