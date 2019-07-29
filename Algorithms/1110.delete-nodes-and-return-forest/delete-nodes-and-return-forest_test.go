package problem1110

import (
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root   []int
	delete []int
	ans    [][]int
}{

	{
		[]int{1, 2, 3, 4, 5, 6, 7},
		[]int{3, 5},
		[][]int{{6}, {7}, {1, 2, kit.NULL, 4}},
	},

	// 可以有多个 testcase
}

func Test_delNodes(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		root := kit.Ints2TreeNode(tc.root)
		tmps := delNodes(root, tc.delete)
		ans := make([][]int, len(tmps))
		for i, t := range tmps {
			ans[i] = kit.Tree2ints(t)
		}
		ast.Equal(tc.ans, ans, "输入:%v", tc)
	}
}

func Benchmark_delNodes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.Ints2TreeNode(tc.root)
			delNodes(root, tc.delete)
		}
	}
}
