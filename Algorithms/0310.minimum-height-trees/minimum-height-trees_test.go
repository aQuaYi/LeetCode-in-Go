package Problem0310

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n     int
	edges [][]int
	ans   []int
}{

	{
		6,
		[][]int{[]int{0, 3}, []int{1, 3}, []int{2, 3}, []int{4, 3}, []int{5, 4}},
		[]int{3, 4},
	},

	{
		4,
		[][]int{[]int{1, 0}, []int{1, 2}, []int{1, 3}},
		[]int{1},
	},

	// 可以有多个 testcase
}

func Test_findMinHeightTrees(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findMinHeightTrees(tc.n, tc.edges), "输入:%v", tc)
	}
}

func Benchmark_findMinHeightTrees(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findMinHeightTrees(tc.n, tc.edges)
		}
	}
}
