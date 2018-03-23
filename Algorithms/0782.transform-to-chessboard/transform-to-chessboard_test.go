package problem0782

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
		[][]int{{0, 0, 0}, {0, 0, 0}, {1, 1, 1}},
		-1,
	},

	{
		[][]int{{0, 0, 1}, {1, 1, 0}, {1, 1, 0}},
		2,
	},

	{
		[][]int{{0, 1, 1}, {0, 1, 1}, {1, 0, 0}},
		2,
	},

	{
		[][]int{{1, 1, 0}, {0, 0, 1}, {0, 0, 1}},
		2,
	},

	{
		[][]int{{0, 1}, {1, 0}},
		0,
	},

	{
		[][]int{{1, 1, 1, 0}, {0, 1, 1, 0}, {1, 0, 0, 1}, {1, 0, 0, 1}},
		-1,
	},

	{
		[][]int{{0, 1, 1, 0}, {0, 1, 1, 0}, {1, 0, 0, 1}, {1, 0, 0, 1}},
		2,
	},

	{
		[][]int{{1, 0}, {1, 0}},
		-1,
	},

	// 可以有多个 testcase
}

func Test_movesToChessboard(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, movesToChessboard(tc.board), "输入:%v", tc)
	}
}

func Benchmark_movesToChessboard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			movesToChessboard(tc.board)
		}
	}
}
