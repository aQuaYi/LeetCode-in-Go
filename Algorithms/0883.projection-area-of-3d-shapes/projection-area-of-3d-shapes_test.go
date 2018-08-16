package problem0883

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	grid [][]int
	ans  int
}{

	{
		[][]int{{2}},
		5,
	},

	{
		[][]int{{1, 2}, {3, 4}},
		17,
	},

	{
		[][]int{{1, 0}, {0, 2}},
		8,
	},

	{
		[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
		14,
	},

	{
		[][]int{{2, 2, 2}, {2, 1, 2}, {2, 2, 2}},
		21,
	},

	// 可以有多个 testcase
}

func Test_projectionArea(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, projectionArea(tc.grid), "输入:%v", tc)
	}
}

func Benchmark_projectionArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			projectionArea(tc.grid)
		}
	}
}
