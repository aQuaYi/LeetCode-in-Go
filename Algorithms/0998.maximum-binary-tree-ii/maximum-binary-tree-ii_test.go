package problem0998

import (
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root []int
	val  int
	ans  []int
}{

	{
		[]int{4, 1, 3, kit.NULL, kit.NULL, 2},
		5,
		[]int{5, 4, kit.NULL, 1, 3, kit.NULL, kit.NULL, 2},
	},

	{
		[]int{5, 2, 4, kit.NULL, 1},
		3,
		[]int{5, 2, 4, kit.NULL, 1, kit.NULL, 3},
	},

	{
		[]int{5, 2, 3, kit.NULL, 1},
		4,
		[]int{5, 2, 4, kit.NULL, 1, 3},
	},

	// 可以有多个 testcase
}

func Test_insertIntoMaxTree(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		root := kit.Ints2TreeNode(tc.root)
		ans := kit.Ints2TreeNode(tc.ans)
		ast.Equal(ans, insertIntoMaxTree(root, tc.val), "输入:%v", tc)
	}
}

func Benchmark_insertIntoMaxTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.Ints2TreeNode(tc.root)
			insertIntoMaxTree(root, tc.val)
		}
	}
}
