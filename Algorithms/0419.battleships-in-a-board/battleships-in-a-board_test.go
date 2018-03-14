package problem0419

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	board [][]byte
	ans   int
}{

	{
		[][]byte{},
		0,
	},

	{
		[][]byte{
			[]byte("X.X.X"),
		},
		3,
	},

	{
		[][]byte{
			[]byte("X..X"),
			[]byte("...X"),
			[]byte("X..X"),
		},
		3,
	},

	{
		[][]byte{
			[]byte("X..X"),
			[]byte("...X"),
			[]byte("...X"),
		},
		2,
	},

	// 可以有多个 testcase
}

func Test_countBattleships(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, countBattleships(tc.board), "输入:%v", tc)
	}
}

func Benchmark_countBattleships(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			countBattleships(tc.board)
		}
	}
}
