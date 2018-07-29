package problem0865

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root []int
	ans  []int
}{

	{
		[]int{3, 5, 1, 6, 2, 0, 8, kit.NULL, kit.NULL, 7, 4},
		[]int{2, 7, 4},
	},

	{
		[]int{3, 5, 1, 6, 2, 0, 8, kit.NULL, kit.NULL, 7, 4, 9, 10},
		[]int{3, 5, 1, 6, 2, 0, 8, kit.NULL, kit.NULL, 7, 4, 9, 10},
	},

	// 可以有多个 testcase
}

func Test_subtreeWithAllDeepest(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := kit.Ints2TreeNode(tc.root)
		actual := kit.Tree2Preorder(subtreeWithAllDeepest(root))
		expected := kit.Tree2Preorder(kit.Ints2TreeNode(tc.ans))
		ast.Equal(expected, actual, "输入:%v", tc)
	}
}

func Benchmark_subtreeWithAllDeepest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.Ints2TreeNode(tc.root)
			subtreeWithAllDeepest(root)
		}
	}
}
