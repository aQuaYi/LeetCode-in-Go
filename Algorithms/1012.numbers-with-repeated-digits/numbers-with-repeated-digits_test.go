package problem1012

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	N   int
	ans int
}{

	{
		1000000000,
		994388230,
	},

	{
		888,
		230,
	},

	{
		20,
		1,
	},

	{
		100,
		10,
	},

	{
		1000,
		262,
	},

	// 可以有多个 testcase
}

func Test_numDupDigitsAtMostN(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, numDupDigitsAtMostN(tc.N), "输入:%v", tc)
	}
}

func Benchmark_numDupDigitsAtMostN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numDupDigitsAtMostN(tc.N)
		}
	}
}
