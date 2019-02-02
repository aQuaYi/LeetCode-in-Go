package problem0935

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
		1,
		10,
	},

	{
		2,
		20,
	},

	{
		3,
		46,
	},

	{
		5000,
		406880451,
	},

	// 可以有多个 testcase
}

func Test_knightDialer(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, knightDialer(tc.N), "输入:%v", tc)
	}
}

func Benchmark_knightDialer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			knightDialer(tc.N)
		}
	}
}
