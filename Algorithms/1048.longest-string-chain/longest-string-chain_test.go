package problem1048

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	words []string
	ans   int
}{

	{
		[]string{"ksqvsyq", "ks", "kss", "czvh", "zczpzvdhx", "zczpzvh", "zczpzvhx", "zcpzvh", "zczvh", "gr", "grukmj", "ksqvsq", "gruj", "kssq", "ksqsq", "grukkmj", "grukj", "zczpzfvdhx", "gru"},
		7,
	},

	{
		[]string{"a", "b", "ba", "bca", "bda", "bdca"},
		4,
	},

	// 可以有多个 testcase
}

func Test_longestStrChain(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, longestStrChain(tc.words), "输入:%v", tc)
	}
}

func Benchmark_longestStrChain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			longestStrChain(tc.words)
		}
	}
}
