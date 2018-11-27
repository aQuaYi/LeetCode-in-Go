package problem0913

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	graph [][]int
	ans   int
}{

	{
		[][]int{
			{2, 5}, {3}, {0, 4, 5}, {1, 4, 5}, {2, 3}, {0, 2, 3},
		},
		0,
	},

	{
		[][]int{
			{6}, {4}, {9}, {5}, {1, 5}, {3, 4, 6}, {0, 5, 10}, {8, 9, 10}, {7}, {2, 7}, {6, 7},
		},
		1,
	},

	// 可以有多个 testcase
}

func Test_catMouseGame(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, catMouseGame(tc.graph), "输入:%v", tc)
	}
}

func Benchmark_catMouseGame(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			catMouseGame(tc.graph)
		}
	}
}
