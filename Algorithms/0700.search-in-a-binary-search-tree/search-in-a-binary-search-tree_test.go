package problem0700

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
		[]int{4, 2, 7, 1, 3},
		6,
		nil,
	},

	{
		[]int{4, 2, 7, 1, 3},
		2,
		[]int{2, 1, 3},
	},

	// 可以有多个 testcase
}

func Test_searchBST(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := kit.Ints2TreeNode(tc.root)
		ans := kit.Tree2Preorder(searchBST(root, tc.val))
		ast.Equal(tc.ans, ans, "输入:%v", tc)
	}
}

func Benchmark_searchBST(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.Ints2TreeNode(tc.root)
			searchBST(root, tc.val)
		}
	}
}
