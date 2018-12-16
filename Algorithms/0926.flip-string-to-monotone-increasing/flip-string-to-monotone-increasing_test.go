package problem0926

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	S   string
	ans int
}{

	{
		"01000011110",
		2,
	},

	{
		"00110",
		1,
	},

	{
		"00110",
		1,
	},

	{
		"010110",
		2,
	},

	{
		"00011000",
		2,
	},

	// 可以有多个 testcase
}

func Test_minFlipsMonoIncr(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, minFlipsMonoIncr(tc.S), "输入:%v", tc)
	}
}

func Benchmark_minFlipsMonoIncr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minFlipsMonoIncr(tc.S)
		}
	}
}
