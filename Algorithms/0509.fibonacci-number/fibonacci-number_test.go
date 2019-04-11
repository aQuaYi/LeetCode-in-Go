package problem0509

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
		30,
		832040,
	},

	{
		2,
		1,
	},

	{
		3,
		2,
	},

	{
		0,
		0,
	},

	{
		4,
		3,
	},

	// 可以有多个 testcase
}

func Test_fib(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, fib(tc.N), "输入:%v", tc)
	}
}

func Benchmark_fib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			fib(tc.N)
		}
	}
}
