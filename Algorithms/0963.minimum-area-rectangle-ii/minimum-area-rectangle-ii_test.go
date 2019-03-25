package problem0963

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	points [][]int
	ans    float64
}{

	{
		[][]int{{1, 2}, {2, 1}, {1, 0}, {0, 1}},
		2.00000,
	},

	{
		[][]int{{0, 1}, {2, 1}, {1, 1}, {1, 0}, {2, 0}},
		1.00000,
	},

	{
		[][]int{{0, 3}, {1, 2}, {3, 1}, {1, 3}, {2, 1}},
		0,
	},

	{
		[][]int{{3, 1}, {1, 1}, {0, 1}, {2, 1}, {3, 3}, {3, 2}, {0, 2}, {2, 3}},
		2.00000,
	},

	// 可以有多个 testcase
}

func Test_minAreaFreeRect(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.InDelta(tc.ans, minAreaFreeRect(tc.points), 0.00001, "输入:%v", tc)
	}
}

func Benchmark_minAreaFreeRect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minAreaFreeRect(tc.points)
		}
	}
}
