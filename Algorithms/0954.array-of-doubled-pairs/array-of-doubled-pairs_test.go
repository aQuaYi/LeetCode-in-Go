package problem0954

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	ans bool
}{

	{
		[]int{3, 1, 3, 6},
		false,
	},

	{
		[]int{2, 1, 2, 6},
		false,
	},

	{
		[]int{4, -2, 2, -4},
		true,
	},

	{
		[]int{1, 2, 4, 16, 8, 4},
		false,
	},

	// 可以有多个 testcase
}

func Test_canReorderDoubled(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, canReorderDoubled(tc.A), "输入:%v", tc)
	}
}

func Benchmark_canReorderDoubled(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			canReorderDoubled(tc.A)
		}
	}
}
