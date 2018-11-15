package problem0909

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
		[][]int{
			{-1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1},
			{-1, 35, -1, -1, 13, -1},
			{-1, -1, -1, -1, -1, -1},
			{-1, 15, -1, -1, -1, -1}},
		4,
	},

	// 可以有多个 testcase
}

func Test_snakesAndLadders(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, snakesAndLadders(tc.board), "输入:%v", tc)
	}
}

func Benchmark_snakesAndLadders(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			snakesAndLadders(tc.board)
		}
	}
}
func Test_location(t *testing.T) {
	ast := assert.New(t)
	//
	const (
		m = 3
		n = 3
	)
	//
	expected := [m][n]int{
		{7, 8, 9},
		{6, 5, 4},
		{1, 2, 3},
	}
	actual := [m][n]int{}
	for i := 1; i <= 9; i++ {
		x, y := location(m, n, i)
		actual[x][y] = i
	}
	ast.Equal(expected, actual)
}
