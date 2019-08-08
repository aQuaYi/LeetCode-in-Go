package problem1137

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n   int
	ans int
}{

	{
		35,
		615693474,
	},

	{
		0,
		0,
	},

	{
		4,
		4,
	},

	{
		25,
		1389537,
	},

	// 可以有多个 testcase
}

func Test_tribonacci(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, tribonacci(tc.n), "输入:%v", tc)
	}
}

func Benchmark_tribonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			tribonacci(tc.n)
		}
	}
}
