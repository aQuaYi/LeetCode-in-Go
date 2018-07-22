package problem0701

import (
	"fmt"
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
		nil,
		4,
		[]int{4},
	},

	{
		[]int{4, 2, 7, 1, 3},
		5,
		[]int{4, 2, 1, 3, 7, 5}, // preorder
	},

	// 可以有多个 testcase
}

func Test_insertIntoBST(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := kit.Ints2TreeNode(tc.root)
		ans := kit.Tree2Preorder(insertIntoBST(root, tc.val))
		ast.Equal(tc.ans, ans, "输入:%v", tc)
	}
}

func Benchmark_insertIntoBST(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.Ints2TreeNode(tc.root)
			insertIntoBST(root, tc.val)
		}
	}
}
