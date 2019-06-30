package problem1041

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	instructions string
	ans          bool
}{

	{
		"GLRLLGLL",
		true,
	},

	{
		"GLGLGGLGL",
		false,
	},

	{
		"LL",
		true,
	},

	{
		"GGLLGG",
		true,
	},

	{
		"GG",
		false,
	},

	{
		"GL",
		true,
	},

	// 可以有多个 testcase
}

func Test_isRobotBounded(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, isRobotBounded(tc.instructions), "输入:%v", tc)
	}
}

func Benchmark_isRobotBounded(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isRobotBounded(tc.instructions)
		}
	}
}
