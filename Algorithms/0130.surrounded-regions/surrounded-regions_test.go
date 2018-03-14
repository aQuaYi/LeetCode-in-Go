package problem0130

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	board [][]byte
	ans   [][]byte
}{

	{
		[][]byte{
			[]byte("OXXXXXOO"),
			[]byte("OOOXXXXO"),
			[]byte("XXXXOOOO"),
			[]byte("XOXOOXXX"),
			[]byte("OXOXXXOO"),
			[]byte("OXXOOXXO"),
			[]byte("OXOXXXOO"),
			[]byte("OXXXXOXX"),
		},
		[][]byte{
			[]byte("OXXXXXOO"),
			[]byte("OOOXXXXO"),
			[]byte("XXXXOOOO"),
			[]byte("XXXOOXXX"),
			[]byte("OXXXXXOO"),
			[]byte("OXXXXXXO"),
			[]byte("OXXXXXOO"),
			[]byte("OXXXXOXX"),
		},
	},

	{
		[][]byte{
			[]byte("XXXXOX"),
			[]byte("OXXOOX"),
			[]byte("XOXOOO"),
			[]byte("XOOOXO"),
			[]byte("OOXXOX"),
			[]byte("XOXOXX"),
		},
		[][]byte{
			[]byte("XXXXOX"),
			[]byte("OXXOOX"),
			[]byte("XOXOOO"),
			[]byte("XOOOXO"),
			[]byte("OOXXXX"),
			[]byte("XOXOXX"),
		},
	},

	{
		[][]byte{
			[]byte("OXOOXX"),
			[]byte("OXXXOX"),
			[]byte("XOOXOO"),
			[]byte("XOXXXX"),
			[]byte("OOXOXX"),
			[]byte("XXOOOO"),
		},
		[][]byte{
			[]byte("OXOOXX"),
			[]byte("OXXXOX"),
			[]byte("XOOXOO"),
			[]byte("XOXXXX"),
			[]byte("OOXOXX"),
			[]byte("XXOOOO"),
		},
	},

	{
		[][]byte{
			[]byte("XOXOXOOOXO"),
			[]byte("XOOXXXOOOX"),
			[]byte("OOOOOOOOXX"),
			[]byte("OOOOOOXOOX"),
			[]byte("OOXXOXXOOO"),
			[]byte("XOOXXXOXXO"),
			[]byte("XOXOOXXOXO"),
			[]byte("XXOXXOXOOX"),
			[]byte("OOOOXOXOXO"),
			[]byte("XXOXXXXOOO"),
		},
		[][]byte{
			[]byte("XOXOXOOOXO"),
			[]byte("XOOXXXOOOX"),
			[]byte("OOOOOOOOXX"),
			[]byte("OOOOOOXOOX"),
			[]byte("OOXXOXXOOO"),
			[]byte("XOOXXXXXXO"),
			[]byte("XOXXXXXOXO"),
			[]byte("XXOXXXXOOX"),
			[]byte("OOOOXXXOXO"),
			[]byte("XXOXXXXOOO"),
		},
	},

	{
		[][]byte{
			[]byte("XXXX"),
			[]byte("XOOX"),
			[]byte("XOOX"),
			[]byte("XXXX"),
		},
		[][]byte{
			[]byte("XXXX"),
			[]byte("XXXX"),
			[]byte("XXXX"),
			[]byte("XXXX"),
		},
	},

	{
		[][]byte{
			[]byte("OXXOX"),
			[]byte("XOOXO"),
			[]byte("XOXOX"),
			[]byte("OXOOO"),
			[]byte("XXOXO"),
		},
		[][]byte{
			[]byte("OXXOX"),
			[]byte("XXXXO"),
			[]byte("XXXOX"),
			[]byte("OXOOO"),
			[]byte("XXOXO"),
		},
	},

	{
		[][]byte{
			[]byte("XXXX"),
			[]byte("XOOX"),
			[]byte("XXOX"),
			[]byte("XOXX"),
		},
		[][]byte{
			[]byte("XXXX"),
			[]byte("XXXX"),
			[]byte("XXXX"),
			[]byte("XOXX"),
		},
	},

	{
		[][]byte{
			[]byte("XXXX"),
		},
		[][]byte{
			[]byte("XXXX"),
		},
	},

	{
		[][]byte{
			[]byte("X"),
			[]byte("X"),
			[]byte("X"),
		},
		[][]byte{
			[]byte("X"),
			[]byte("X"),
			[]byte("X"),
		},
	},

	// 可以有多个 testcase
}

func Test_solve(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		solve(tc.board)
		ast.Equal(tc.ans, tc.board, "输入:%v", tc)
	}
}

func Benchmark_solve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			solve(tc.board)
		}
	}
}
