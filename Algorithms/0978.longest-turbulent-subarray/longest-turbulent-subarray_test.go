package problem0978

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
		[]int{9, 4, 2, 10, 7, 8, 8, 1, 9},
		5,
	},

	{
		[]int{4, 8, 12, 16},
		2,
	},

	{
		[]int{100},
		1,
	},

	// 可以有多个 testcase
}

func Test_maxTurbulenceSize(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, maxTurbulenceSize(tc.A), "输入:%v", tc)
	}
}

func Benchmark_maxTurbulenceSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxTurbulenceSize(tc.A)
		}
	}
}
