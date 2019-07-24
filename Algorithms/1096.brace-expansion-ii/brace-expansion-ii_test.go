package problem1096

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	expression string
	ans        []string
}{

	{
		"{a,b}{c,{d,e}}",
		[]string{"ac", "ad", "ae", "bc", "bd", "be"},
	},

	{
		"{{a,z},a{b,c},{ab,z}}",
		[]string{"a", "ab", "ac", "z"},
	},

	// 可以有多个 testcase
}

func Test_braceExpansionII(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, braceExpansionII(tc.expression), "输入:%v", tc)
	}
}

func Benchmark_braceExpansionII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			braceExpansionII(tc.expression)
		}
	}
}
