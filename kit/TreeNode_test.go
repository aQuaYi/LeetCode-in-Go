package kit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	preOrder  = []int{1, 2, 4, 5, 3, 6, 7}
	inOrder   = []int{4, 2, 5, 1, 6, 3, 7}
	postOrder = []int{4, 5, 2, 6, 7, 3, 1}
)

func Test_preIn2Tree(t *testing.T) {
	ast := assert.New(t)

	actual := Tree2Postorder(PreIn2Tree(preOrder, inOrder))
	expected := postOrder
	ast.Equal(expected, actual)

	ast.Panics(func() { PreIn2Tree([]int{1}, []int{}) })

	ast.Nil(PreIn2Tree([]int{}, []int{}))
}

func Test_inPost2Tree(t *testing.T) {
	ast := assert.New(t)

	actual := Tree2Preorder(InPost2Tree(inOrder, postOrder))
	expected := preOrder
	ast.Equal(expected, actual)

	ast.Panics(func() { InPost2Tree([]int{1}, []int{}) })

	ast.Nil(InPost2Tree([]int{}, []int{}))
}

func Test_tree2Ints(t *testing.T) {
	ast := assert.New(t)
	root := PreIn2Tree(preOrder, inOrder)

	ast.Equal(preOrder, Tree2Preorder(root))
	ast.Nil(Tree2Preorder(nil))

	ast.Equal(inOrder, Tree2Inorder(root))
	ast.Nil(Tree2Inorder(nil))

	ast.Equal(postOrder, Tree2Postorder(root))
	ast.Nil(Tree2Postorder(nil))
}

func Test_indexOf(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(1, indexOf(1, []int{0, 1, 2, 3}))

	ast.Panics(func() { indexOf(0, []int{1, 2, 3}) })
}
