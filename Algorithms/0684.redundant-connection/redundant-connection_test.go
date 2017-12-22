package Problem0684

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
		[][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}, {1, 5}},
		[]int{1, 5},
	},

	{
		[][]int{{1, 3}, {3, 4}, {1, 5}, {3, 5}, {2, 3}},
		[]int{3, 5},
	},

	{
		[][]int{{1, 4}, {3, 4}, {1, 3}, {1, 2}, {4, 5}},
		[]int{1, 3},
	},

	{
		[][]int{{1, 2}, {1, 3}, {2, 3}},
		[]int{2, 3},
	},

	{
		[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}},
		[]int{1, 4},
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findRedundantConnection(tc.edges), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findRedundantConnection(tc.edges)
		}
	}
}
