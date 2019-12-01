package problem0337

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  int
}{

	// 用 0 表示 nil
	{
		[]int{4, 1, 0, 2, 0, 3},
		7,
	},

	{
		[]int{3, 2, 3, 0, 3, 0, 1},
		7,
	},

	{
		[]int{3, 4, 5, 1, 3, 0, 1},
		9,
	},

	// 可以有多个 testcase
}

func Test_rob(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := &TreeNode{}
		initTree(tc.nums, []*TreeNode{}, root)
		ast.Equal(tc.ans, rob(root), "输入:%v", tc)
	}
}

func Benchmark_rob(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := &TreeNode{}
			initTree(tc.nums, []*TreeNode{}, root)
			rob(root)
		}
	}
}

// 创建二叉树
func initTree(nums []int, parents []*TreeNode, root *TreeNode) {
	if len(nums) == 0 {
		return
	}

	if len(parents) == 0 {
		if nums[0] != 0 {
			root.Val = nums[0]
			parents = append(parents, root)
			initTree(nums[1:], parents, root)
		}

		return
	}

	newParents := make([]*TreeNode, 0, len(parents)*2)
	for _, parent := range parents {

		if nums[0] != 0 {
			parent.Left = &TreeNode{Val: nums[0]}
			newParents = append(newParents, parent.Left)
		}

		if len(nums) == 1 {
			return
		}

		if nums[1] != 0 {
			parent.Right = &TreeNode{Val: nums[1]}
			newParents = append(newParents, parent.Right)
		}

		nums = nums[2:]
	}

	initTree(nums, newParents, root)
}
