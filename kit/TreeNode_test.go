package kit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	preOrder  = []int{1, 2, 4, 5, 3, 6, 7}
	inOrder   = []int{4, 2, 5, 1, 6, 3, 7}
	postOrder = []int{4, 5, 2, 6, 7, 3, 1}
)

func Test_preIn2Tree(t *testing.T) {
	ast := assert.New(t)

	actual := tree2postorder(preIn2Tree(preOrder, inOrder))
	expected := postOrder
	ast.Equal(expected, actual)

	ast.Panics(func() { preIn2Tree([]int{1}, []int{}) })

	ast.Nil(preIn2Tree([]int{}, []int{}))
}

func Test_inPost2Tree(t *testing.T) {
	ast := assert.New(t)

	actual := tree2preorder(inPost2Tree(inOrder, postOrder))
	expected := preOrder
	ast.Equal(expected, actual)

	ast.Panics(func() { inPost2Tree([]int{1}, []int{}) })

	ast.Nil(inPost2Tree([]int{}, []int{}))
}

func Test_tree2Ints(t *testing.T) {
	ast := assert.New(t)
	root := preIn2Tree(preOrder, inOrder)

	ast.Equal(preOrder, tree2preorder(root))
	ast.Equal(inOrder, tree2inorder(root))
	ast.Equal(postOrder, tree2postorder(root))
}

func Test_indexOf(t *testing.T) {
	ast := assert.New(t)

	ast.Panics(func() { indexOf(0, []int{1, 2, 3}) })
}
