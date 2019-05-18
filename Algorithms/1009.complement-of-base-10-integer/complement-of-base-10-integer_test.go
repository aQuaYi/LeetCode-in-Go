package problem1009

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
		2,
		1,
	},

	{
		1,
		0,
	},

	{
		0,
		1,
	},

	{
		5,
		2,
	},

	{
		7,
		0,
	},

	{
		10,
		5,
	},

	{
		87654321,
		46563406,
	},

	// 可以有多个 testcase
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, bitwiseComplement(tc.N), "输入:%v", tc)
	}
}

func Benchmark_bitwiseComplement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			bitwiseComplement(tc.N)
		}
	}
}
