package problem0290

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	pattern string
	str     string
	ans     bool
}{

	{"abba", "dog cat cat dog", true},
	{"abba", "dog cat cat fish", false},
	{"aaaa", "dog cat cat dog", false},
	{"abba", "dog dog dog dog", false},
	{"aaa", "aa aa aa aa", false},

	// 可以有多个 testcase
}

func Test_wordPattern(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, wordPattern(tc.pattern, tc.str), "输入:%v", tc)
	}
}

func Benchmark_wordPattern(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			wordPattern(tc.pattern, tc.str)
		}
	}
}
