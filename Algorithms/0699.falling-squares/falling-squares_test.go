package Problem0699

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
