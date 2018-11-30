package problem0916

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A, B []string
	ans  []string
}{

	{
		[]string{"amazon", "apple", "facebook", "google", "leetcode"},
		[]string{"e", "o"},
		[]string{"facebook", "google", "leetcode"},
	},

	{
		[]string{"amazon", "apple", "facebook", "google", "leetcode"},
		[]string{"l", "e"},
		[]string{"apple", "google", "leetcode"},
	},

	{
		[]string{"amazon", "apple", "facebook", "google", "leetcode"},
		[]string{"e", "oo"},
		[]string{"facebook", "google"},
	},

	{
		[]string{"amazon", "apple", "facebook", "google", "leetcode"},
		[]string{"lo", "eo"},
		[]string{"google", "leetcode"},
	},

	{
		[]string{"amazon", "apple", "facebook", "google", "leetcode"},
		[]string{"ec", "oc", "ceo"},
		[]string{"facebook", "leetcode"},
	},

	// 可以有多个 testcase
}

func Test_wordSubsets(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, wordSubsets(tc.A, tc.B), "输入:%v", tc)
	}
}

func Benchmark_wordSubsets(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			wordSubsets(tc.A, tc.B)
		}
	}
}
