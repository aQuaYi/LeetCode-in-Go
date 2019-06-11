package problem1015

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	K   int
	ans int
}{

	{
		23,
		22,
	},

	{
		5,
		-1,
	},

	{
		1,
		1,
	},

	{
		2,
		-1,
	},

	{
		3,
		3,
	},

	// 可以有多个 testcase
}

func Test_smallestRepunitDivByK(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, smallestRepunitDivByK(tc.K), "输入:%v", tc)
	}
}

func Benchmark_smallestRepunitDivByK(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			smallestRepunitDivByK(tc.K)
		}
	}
}
