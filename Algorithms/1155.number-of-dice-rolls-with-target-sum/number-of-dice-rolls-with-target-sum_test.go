package problem1155

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	d      int
	f      int
	target int
	ans    int
}{

	{
		1,
		6,
		3,
		1,
	},

	{
		2,
		6,
		7,
		6,
	},

	{
		2,
		5,
		10,
		1,
	},

	{
		1,
		2,
		3,
		0,
	},

	{
		30,
		30,
		500,
		222616187,
	},

	// 可以有多个 testcase
}

func Test_numRollsToTarget(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, numRollsToTarget(tc.d, tc.f, tc.target), "输入:%v", tc)
	}
}

func Benchmark_numRollsToTarget(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numRollsToTarget(tc.d, tc.f, tc.target)
		}
	}
}
