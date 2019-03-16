package problem0953

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	words []string
	order string
	ans   bool
}{

	{
		[]string{"hello", "leetcode"},
		"hlabcdefgijkmnopqrstuvwxyz",
		true,
	},

	{
		[]string{"word", "world", "row"},
		"worldabcefghijkmnpqstuvxyz",
		false,
	},

	{
		[]string{"apple", "app"},
		"abcdefghijklmnopqrstuvwxyz",
		false,
	},

	// 可以有多个 testcase
}

func Test_isAlienSorted(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, isAlienSorted(tc.words, tc.order), "输入:%v", tc)
	}
}

func Benchmark_isAlienSorted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isAlienSorted(tc.words, tc.order)
		}
	}
}
