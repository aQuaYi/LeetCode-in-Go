package problem0794

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	board []string
	ans   bool
}{

	{
		[]string{"O  ", "   ", "   "},
		false,
	},

	{
		[]string{"XOX", " X ", "   "},
		false,
	},

	{
		[]string{"X X", "X  ", "OOO"},
		true,
	},

	{
		[]string{"XXX", "   ", "OOO"},
		false,
	},

	{
		[]string{"XOX", "O O", "XOX"},
		true,
	},

	// 可以有多个 testcase
}

func Test_validTicTacToe(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, validTicTacToe(tc.board), "输入:%v", tc)
	}
}

func Benchmark_validTicTacToe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			validTicTacToe(tc.board)
		}
	}
}
