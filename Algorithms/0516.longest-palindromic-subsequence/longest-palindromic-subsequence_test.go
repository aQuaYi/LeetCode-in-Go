package problem0516

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
		"bbbab",
		4,
	},

	{
		"cbbd",
		2,
	},

	// 可以有多个 testcase
}

func Test_longestPalindromeSubseq(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, longestPalindromeSubseq(tc.s), "输入:%v", tc)
	}
}

func Benchmark_longestPalindromeSubseq(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			longestPalindromeSubseq(tc.s)
		}
	}
}
