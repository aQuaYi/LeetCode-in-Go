package problem1016

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	S   string
	N   int
	ans bool
}{

	{
		"0110",
		3,
		true,
	},

	{
		"01010",
		4,
		false,
	},

	{
		"0110",
		4,
		false,
	},

	// 可以有多个 testcase
}

func Test_queryString(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, queryString(tc.S, tc.N), "输入:%v", tc)
	}
}

func Benchmark_queryString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			queryString(tc.S, tc.N)
		}
	}
}
