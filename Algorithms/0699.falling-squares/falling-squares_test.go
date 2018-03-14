package problem0699

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	positions [][]int
	ans       []int
}{

	{
		[][]int{{33, 34}, {47, 62}, {70, 16}, {90, 79}, {73, 86}, {55, 6}, {74, 2}, {40, 95}, {52, 16}, {50, 33}},
		[]int{34, 96, 112, 175, 261, 261, 263, 358, 374, 407},
	},

	{
		[][]int{{1, 2}, {2, 3}, {6, 1}},
		[]int{2, 5, 5},
	},

	{
		[][]int{{100, 100}, {200, 100}},
		[]int{100, 100},
	},

	{
		[][]int{{1, 2}, {3, 3}, {2, 3}, {6, 1}},
		[]int{2, 3, 6, 6},
	},

	{
		[][]int{{1, 2}, {4, 3}, {2, 4}, {6, 1}},
		[]int{2, 3, 7, 7},
	},

	// 可以有多个 testcase
}

func Test_fallingSquares(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, fallingSquares(tc.positions), "输入:%v", tc)
	}
}

func Benchmark_fallingSquares(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			fallingSquares(tc.positions)
		}
	}
}
