package problem0529

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	board [][]byte
	click []int
	ans   [][]byte
}{

	{
		[][]byte{
			{'E', 'E', 'E', 'E', 'E'},
			{'E', 'E', 'M', 'E', 'E'},
			{'E', 'E', 'E', 'E', 'E'},
			{'E', 'E', 'E', 'E', 'E'}},
		[]int{3, 0},
		[][]byte{
			{'B', '1', 'E', '1', 'B'},
			{'B', '1', 'M', '1', 'B'},
			{'B', '1', '1', '1', 'B'},
			{'B', 'B', 'B', 'B', 'B'}},
	},

	{
		[][]byte{
			{'B', '1', 'E', '1', 'B'},
			{'B', '1', 'M', '1', 'B'},
			{'B', '1', '1', '1', 'B'},
			{'B', 'B', 'B', 'B', 'B'}},
		[]int{1, 2},
		[][]byte{
			{'B', '1', 'E', '1', 'B'},
			{'B', '1', 'X', '1', 'B'},
			{'B', '1', '1', '1', 'B'},
			{'B', 'B', 'B', 'B', 'B'}},
	},

	// 可以有多个 testcase
}

func Test_updateBoard(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, updateBoard(tc.board, tc.click), "输入:%v", tc)
	}
}

func Benchmark_updateBoard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			updateBoard(tc.board, tc.click)
		}
	}
}
