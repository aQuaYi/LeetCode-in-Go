package problem1017

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	N   int
	ans string
}{

	{
		0,
		"0",
	},

	{
		6,
		"11010",
	},

	{
		2,
		"110",
	},

	{
		3,
		"111",
	},

	{
		4,
		"100",
	},

	// 可以有多个 testcase
}

func Test_baseNeg2(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, baseNeg2(tc.N), "输入:%v", tc)
	}
}

func Benchmark_baseNeg2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			baseNeg2(tc.N)
		}
	}
}
