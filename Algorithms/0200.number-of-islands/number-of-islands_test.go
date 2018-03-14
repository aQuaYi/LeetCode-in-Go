package problem0200

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	grid [][]byte
	ans  int
}{

	{
		[][]byte{},
		0,
	},

	{
		[][]byte{
			[]byte("111001"),
			[]byte("010001"),
			[]byte("111111"),
		},
		1,
	},

	{
		[][]byte{
			[]byte("11110"),
			[]byte("11010"),
			[]byte("11000"),
			[]byte("00000"),
		},
		1,
	},

	{
		[][]byte{
			[]byte("11000"),
			[]byte("11000"),
			[]byte("00100"),
			[]byte("00011"),
		},
		3,
	},

	// 可以有多个 testcase
}

func Test_numIslands(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, numIslands(tc.grid), "输入:%v", tc)
	}
}

func Benchmark_numIslands(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numIslands(tc.grid)
		}
	}
}
