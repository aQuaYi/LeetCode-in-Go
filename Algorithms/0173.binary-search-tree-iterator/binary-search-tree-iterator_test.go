package problem0173

import (
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

func Test_BSTIterator(t *testing.T) {
	ast := assert.New(t)

	ints := []int{7, 3, 15, kit.NULL, kit.NULL, 9, 20}

	root := kit.Ints2TreeNode(ints)

	it := Constructor(root)

	i := 0

	for it.HasNext() {
		for ints[i] == kit.NULL {
			i++
		}
		ast.Equal(ints[i], it.Next(), "%d", ints[i])
	}
}
