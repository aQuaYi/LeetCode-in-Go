package problem0990

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	equations []string
	ans       bool
}{

	{
		[]string{"a==b", "b!=a"},
		false,
	},

	{
		[]string{"b==a", "a==b"},
		true,
	},

	{
		[]string{"a==b", "b==c", "a==c"},
		true,
	},

	{
		[]string{"b!=b"},
		false,
	},

	{
		[]string{"a==b", "b!=c", "c==a"},
		false,
	},

	{
		[]string{"c==c", "b==d", "x!=z"},
		true,
	},

	{
		[]string{"a==b", "c!=d"},
		true,
	},

	// 可以有多个 testcase
}

func Test_equationsPossible(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, equationsPossible(tc.equations), "输入:%v", tc)
	}
}

func Benchmark_equationsPossible(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			equationsPossible(tc.equations)
		}
	}
}
