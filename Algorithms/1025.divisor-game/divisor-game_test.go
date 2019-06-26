package problem1025

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	N   int
	ans bool
}{

	{
		1000,
		true,
	},

	{
		2,
		true,
	},

	{
		3,
		false,
	},

	// 可以有多个 testcase
}

func Test_divisorGame(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, divisorGame(tc.N), "输入:%v", tc)
	}
}

func Benchmark_divisorGame(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			divisorGame(tc.N)
		}
	}
}
