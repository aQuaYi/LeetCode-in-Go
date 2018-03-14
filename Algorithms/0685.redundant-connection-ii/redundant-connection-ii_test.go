package problem0685

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	edges [][]int
	ans   []int
}{

	{
		[][]int{{2, 1}, {3, 1}, {4, 2}, {1, 4}},
		[]int{2, 1},
	},

	{
		[][]int{{1, 2}, {1, 3}, {2, 3}},
		[]int{2, 3},
	},

	{
		[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 5}},
		[]int{4, 1},
	},

	// 可以有多个 testcase
}

func Test_findRedundantDirectedConnection(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findRedundantDirectedConnection(tc.edges), "输入:%v", tc)
	}
}

func Benchmark_findRedundantDirectedConnection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findRedundantDirectedConnection(tc.edges)
		}
	}
}
