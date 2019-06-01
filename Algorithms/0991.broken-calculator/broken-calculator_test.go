package problem0991

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	X   int
	Y   int
	ans int
}{

	{
		68,
		71,
		34,
	},

	{
		1,
		1000000000,
		39,
	},

	{
		2,
		3,
		2,
	},

	{
		5,
		8,
		2,
	},

	{
		3,
		10,
		3,
	},

	{
		1024,
		1,
		1023,
	},

	// 可以有多个 testcase
}

func Test_brokenCalc(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, brokenCalc(tc.X, tc.Y), "输入:%v", tc)
	}
}

func Benchmark_brokenCalc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			brokenCalc(tc.X, tc.Y)
		}
	}
}
