package problem0773

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	board [][]int
	ans   int
}{

	{
		[][]int{{4, 1, 2}, {5, 0, 3}},
		5,
	},

	{
		[][]int{{1, 2, 3}, {4, 0, 5}},
		1,
	},

	{
		[][]int{{3, 2, 4}, {1, 5, 0}},
		14,
	},

	{
		[][]int{{1, 2, 3}, {4, 5, 0}},
		0,
	},

	{
		[][]int{{1, 2, 3}, {5, 4, 0}},
		-1,
	},

	// 可以有多个 testcase
}

func Test_slidingPuzzle(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, slidingPuzzle(tc.board), "输入:%v", tc)
	}
}

func Benchmark_slidingPuzzle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			slidingPuzzle(tc.board)
		}
	}
}
