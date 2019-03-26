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

	{
		3,
		6,
		1,
	},

	{
		3,
		55125018,
		105,
	},

	{
		30,
		281347,
		81,
	},

	{
		97,
		49,
		96,
	},

	{
		11,
		19,
		7,
	},

	{
		11,
		500041,
		41,
	},

	{
		2,
		6125100,
		104,
	},
	{
		3,
		2,
		2,
	},

	{
		3,
		365,
		17,
	},

	{
		100,
		100000000,
		3,
	},

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
