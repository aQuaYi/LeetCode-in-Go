package problem1006

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
		4,
		7,
	},

	{
		10,
		12,
	},

	// 可以有多个 testcase
}

func Test_clumsy(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, clumsy(tc.N), "输入:%v", tc)
	}
}

func Benchmark_clumsy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			clumsy(tc.N)
		}
	}
}
