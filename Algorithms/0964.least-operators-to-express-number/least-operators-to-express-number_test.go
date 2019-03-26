package problem0964

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	x      int
	target int
	ans    int
}{

	// {
	// 	100,
	// 	100000000,
	// 	3,
	// },

	{
		3,
		19,
		5,
	},

	{
		5,
		501,
		8,
	},

	// 可以有多个 testcase
}

func Test_leastOpsExpressTarget(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, leastOpsExpressTarget(tc.x, tc.target), "输入:%v", tc)
	}
}

func Benchmark_leastOpsExpressTarget(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			leastOpsExpressTarget(tc.x, tc.target)
		}
	}
}
