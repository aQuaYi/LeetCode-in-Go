package problem0840

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
		[][]int{
			{3, 2, 9, 2, 7},
			{6, 1, 8, 4, 2},
			{7, 5, 3, 2, 7},
			{2, 9, 4, 9, 6},
			{4, 3, 8, 2, 5},
		},
		1,
	},

	{
		[][]int{
			{4, 3, 8, 4},
			{9, 5, 1, 9},
			{2, 7, 6, 2},
		},
		1,
	},

	{
		[][]int{
			{1, 8, 6},
			{10, 5, 0},
			{4, 2, 9},
		},
		0,
	},

	// 可以有多个 testcase
}

func Test_numMagicSquaresInside(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, numMagicSquaresInside(tc.grid), "输入:%v", tc)
	}
}

func Benchmark_numMagicSquaresInside(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numMagicSquaresInside(tc.grid)
		}
	}
}
