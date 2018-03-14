package problem0318

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	words []string
	ans   int
}{
	{[]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}, 16},
	{[]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}, 4},
	{[]string{"a", "aa", "aaa", "aaaa"}, 0},

	// 可以有多个 testcase
}

func Test_maxProduct(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, maxProduct(tc.words), "输入:%v", tc)
	}
}

func Benchmark_maxProduct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxProduct(tc.words)
		}
	}
}
