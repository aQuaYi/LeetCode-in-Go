package problem0962

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	ans int
}{

	{
		[]int{6, 0, 8, 2, 1, 5},
		4,
	},

	{
		[]int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1},
		7,
	},

	// 可以有多个 testcase
}

func Test_maxWidthRamp(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, maxWidthRamp(tc.A), "输入:%v", tc)
	}
}

func Benchmark_maxWidthRamp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxWidthRamp(tc.A)
		}
	}
}
