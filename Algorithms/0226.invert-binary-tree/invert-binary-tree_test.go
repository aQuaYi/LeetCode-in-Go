package problem0226

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	pre, in []int
	ans     []int // inorder
}{

	{
		[]int{4, 2, 1, 3, 7, 6, 9},
		[]int{1, 2, 3, 4, 6, 7, 9},
		[]int{9, 7, 6, 4, 3, 2, 1},
	},

	// 可以有多个 testcase
}

func Test_invertTree(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := kit.PreIn2Tree(tc.pre, tc.in)
		ans := invertTree(root)
		ast.Equal(tc.ans, kit.Tree2Inorder(ans), "输入:%v", tc)
	}
}
