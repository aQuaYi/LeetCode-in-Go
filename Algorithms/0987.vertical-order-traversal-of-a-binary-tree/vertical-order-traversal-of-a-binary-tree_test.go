package problem0987

import (
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root []int
	ans  [][]int
}{

	{
		[]int{0, 2, 1, 3, kit.NULL, kit.NULL, kit.NULL, 4, 5, kit.NULL, 7, 6, kit.NULL, 10, 8, 11, 9},
		[][]int{{4, 10, 11}, {3, 6, 7}, {2, 5, 8, 9}, {0}, {1}},
	},

	{
		[]int{3, 9, 20, kit.NULL, kit.NULL, 15, 7},
		[][]int{{9}, {3, 15}, {20}, {7}},
	},

	{
		[]int{1, 2, 3, 4, 5, 6, 7},
		[][]int{{4}, {2}, {1, 5, 6}, {3}, {7}},
	},

	// 可以有多个 testcase
}

func Test_verticalTraversal(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		root := kit.Ints2TreeNode(tc.root)
		ast.Equal(tc.ans, verticalTraversal(root), "输入:%v", tc)
	}
}

func Benchmark_verticalTraversal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.Ints2TreeNode(tc.root)
			verticalTraversal(root)
		}
	}
}
