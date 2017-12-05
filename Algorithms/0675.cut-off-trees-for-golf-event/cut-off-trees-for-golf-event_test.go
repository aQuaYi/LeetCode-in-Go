package Problem0675

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	forest [][]int
	ans    int
}{

	{
		[][]int{{18, 17, 16, 15, 14, 13}, {7, 8, 9, 10, 11, 12}, {6, 5, 4, 3, 2, 1}},
		22,
	},

	{
		[][]int{{9, 8, 7}, {4, 5, 6}, {3, 2, 1}},
		10,
	},

	{
		[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		12,
	},

	{
		[][]int{{1, 2, 3}, {0, 1, 0}, {7, 1, 1}},
		6,
	},

	{
		[][]int{{1, 2, 3}, {0, 1, 0}, {7, 6, 5}},
		8,
	},

	{
		[][]int{{1, 2, 3}, {0, 0, 4}, {7, 6, 5}},
		6,
	},

	{
		[][]int{{1, 2, 3}, {0, 0, 0}, {7, 6, 5}},
		-1,
	},

	{
		[][]int{{2, 3, 4}, {0, 0, 5}, {8, 7, 6}},
		6,
	},

	// 可以有多个 testcase
}

func Test_cutOffTree(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, cutOffTree(tc.forest), "输入:%v", tc)
	}
}

func Benchmark_cutOffTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			cutOffTree(tc.forest)
		}
	}
}
