package problem0919

import (
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_CBTInserter(t *testing.T) {
	Convey("insert [1]", t, func() {
		root := &TreeNode{Val: 1}
		c := Constructor(root)

		parent := c.Insert(2)
		So(parent, ShouldEqual, 1)

		actual := c.Get_root()
		expected := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
		}
		So(actual, ShouldResemble, expected)
	})

	Convey("insert [1,2,3,4,5,6]", t, func() {
		root := kit.Ints2TreeNode([]int{1, 2, 3, 4, 5, 6})
		c := Constructor(root)

		parent := c.Insert(7)
		So(parent, ShouldEqual, 3)

		parent = c.Insert(8)
		So(parent, ShouldEqual, 4)

		actual := c.Get_root()
		expected := kit.Ints2TreeNode([]int{1, 2, 3, 4, 5, 6, 7, 8})
		So(actual, ShouldResemble, expected)
	})

	//
}
