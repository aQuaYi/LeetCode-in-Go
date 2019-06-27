package problem1030

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	R   int
	C   int
	r0  int
	c0  int
	ans [][]int
}{

	{
		1,
		2,
		0,
		0,
		[][]int{{0, 0}, {0, 1}},
	},
	{
		2,
		2,
		0,
		1,
		[][]int{{0, 1}, {0, 0}, {1, 1}, {1, 0}},
	},
	{
		2,
		3,
		1,
		2,
		[][]int{{1, 2}, {0, 2}, {1, 1}, {0, 1}, {1, 0}, {0, 0}},
	},

	// 可以有多个 testcase
}

func Test_allCellsDistOrder(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, allCellsDistOrder(tc.R, tc.C, tc.r0, tc.c0), "输入:%v", tc)
	}
}

func Benchmark_allCellsDistOrder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			allCellsDistOrder(tc.R, tc.C, tc.r0, tc.c0)
		}
	}
}
