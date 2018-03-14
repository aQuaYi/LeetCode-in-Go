package problem0140

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s        string
	wordDict []string
	ans      []string
}{

	{
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		[]string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"},
		[]string{},
	},

	{
		"catsanddog",
		[]string{"cat", "cats", "and", "sand", "dog"},
		[]string{"cat sand dog", "cats and dog"},
	},

	{
		"catsanddog",
		[]string{},
		[]string{},
	},

	// 可以有多个 testcase
}

func Test_wordBreak(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, wordBreak(tc.s, tc.wordDict), "输入:%v", tc)
	}
}

func Benchmark_wordBreak(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			wordBreak(tc.s, tc.wordDict)
		}
	}
}
