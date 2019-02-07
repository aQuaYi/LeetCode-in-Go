package problem0952

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
		[]int{4, 6, 15, 35},
		4,
	},

	{
		[]int{20, 50, 9, 63},
		2,
	},

	{
		[]int{2, 3, 6, 7, 4, 12, 21, 39},
		8,
	},

	// 可以有多个 testcase
}

func Test_largestComponentSize(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, largestComponentSize(tc.A), "输入:%v", tc)
	}
}

func Benchmark_largestComponentSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			largestComponentSize(tc.A)
		}
	}
}
