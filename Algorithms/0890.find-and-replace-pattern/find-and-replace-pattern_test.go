package problem0890

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	words   []string
	pattern string
	ans     []string
}{

	{
		[]string{"abc", "deq", ",ee", "aqq", "dkd", "ccc"},
		"abb",
		[]string{",ee", "aqq"},
	},

	{
		[]string{"abc", "deq", "Mee", "aqq", "dkd", "ccc"},
		"abb",
		[]string{"Mee", "aqq"},
	},

	{
		[]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"},
		"abb",
		[]string{"mee", "aqq"},
	},

	// 可以有多个 testcase
}

func Test_findAndReplacePattern(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findAndReplacePattern(tc.words, tc.pattern), "输入:%v", tc)
	}
}

func Benchmark_findAndReplacePattern(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findAndReplacePattern(tc.words, tc.pattern)
		}
	}
}
