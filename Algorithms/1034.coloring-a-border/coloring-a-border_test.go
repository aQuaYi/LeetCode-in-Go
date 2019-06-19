package problem1034

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	grid  [][]int
	r0    int
	c0    int
	color int
	ans   [][]int
}{

	{
		[][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 3, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}},
		0,
		0,
		3,
		[][]int{
			{3, 3, 3, 3, 3},
			{3, 1, 3, 1, 3},
			{3, 3, 3, 3, 3},
			{3, 1, 3, 1, 3},
			{3, 3, 3, 3, 3},
		},
	},

	{
		[][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 3, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}},
		0,
		0,
		2,
		[][]int{
			{2, 2, 2, 2, 2},
			{2, 1, 2, 1, 2},
			{2, 2, 3, 2, 2},
			{2, 1, 2, 1, 2},
			{2, 2, 2, 2, 2},
		},
	},

	{
		[][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}},
		2,
		2,
		2,
		[][]int{{2, 2, 2, 2, 2}, {2, 1, 1, 1, 2}, {2, 1, 1, 1, 2}, {2, 1, 1, 1, 2}, {2, 2, 2, 2, 2}},
	},

	{
		[][]int{{1, 1}, {1, 2}},
		0,
		0,
		3,
		[][]int{{3, 3}, {3, 2}},
	},

	{
		[][]int{{1, 2, 2}, {2, 3, 2}},
		0,
		1,
		3,
		[][]int{{1, 3, 3}, {2, 3, 3}},
	},

	{
		[][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
		1,
		1,
		2,
		[][]int{{2, 2, 2}, {2, 1, 2}, {2, 2, 2}},
	},

	// 可以有多个 testcase
}

func Test_colorBorder(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, colorBorder(tc.grid, tc.r0, tc.c0, tc.color), "输入:%v", tc)
	}
}

func Benchmark_colorBorder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			colorBorder(tc.grid, tc.r0, tc.c0, tc.color)
		}
	}
}
