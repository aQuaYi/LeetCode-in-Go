package problem0417

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	matrix [][]int
	ans    [][]int
}{

	{
		[][]int{},
		[][]int{},
	},

	{
		[][]int{
			{1, 2, 2, 3, 5},
			{3, 2, 3, 4, 4},
			{2, 4, 5, 3, 1},
			{6, 7, 1, 4, 5},
			{5, 1, 1, 2, 4},
		},
		[][]int{
			{0, 4},
			{1, 3},
			{1, 4},
			{2, 2},
			{3, 0},
			{3, 1},
			{4, 0},
		},
	},

	// 可以有多个 testcase
}

func Test_pacificAtlantic(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, pacificAtlantic(tc.matrix), "输入:%v", tc)
	}
}

func Benchmark_pacificAtlantic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			pacificAtlantic(tc.matrix)
		}
	}
}
