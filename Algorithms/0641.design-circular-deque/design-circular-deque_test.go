package problem0641

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_myFunc(t *testing.T) {
	ast := assert.New(t)

	d := Constructor(3)

	ast.False(d.IsFull(), "检查空数组")
	ast.True(d.IsEmpty(), "检查空数组")

	ast.False(d.DeleteFront(), "从空数组中删除首部元素")
	ast.False(d.DeleteLast(), "从空数组中删除尾部元素")

	ast.Equal(-1, d.GetFront(), "从空数组中获取首部元素")
	ast.Equal(-1, d.GetRear(), "从空数组中获取尾部元素")

	for i := 0; i < 100; i++ {
		ast.True(d.InsertFront(i), "在首部插入 %d", i)
		ast.True(d.InsertFront(i), "在首部插入 %d", i)
		ast.True(d.InsertFront(i), "在首部插入 %d", i)
		ast.True(d.DeleteLast(), "删除尾部元素 %d", i)
		ast.True(d.DeleteLast(), "删除尾部元素 %d", i)
		ast.True(d.DeleteLast(), "删除尾部元素 %d", i)
	}

	for i := 0; i < 100; i++ {
		ast.True(d.InsertLast(i), "在尾部插入 %d", i)
		ast.True(d.InsertLast(i), "在尾部插入 %d", i)
		ast.True(d.InsertLast(i), "在尾部插入 %d", i)
		ast.True(d.DeleteFront(), "删除首部元素 %d", i)
		ast.True(d.DeleteFront(), "删除首部元素 %d", i)
		ast.True(d.DeleteFront(), "删除首部元素 %d", i)
	}

	ast.True(d.InsertLast(1), "往尾部插入 1")
	ast.True(d.InsertLast(2), "往尾部插入 2")
	ast.True(d.InsertFront(3), "往首部插入 3")
	ast.False(d.InsertFront(4), "往首部插入 4")
	ast.False(d.InsertLast(4), "往尾部插入 4")

	ast.Equal(2, d.GetRear(), "获取尾部元素")
	ast.Equal(3, d.GetFront(), "获取首部元素")

	ast.True(d.IsFull(), "检查满数组")
	ast.False(d.IsEmpty(), "检查满数组")

	ast.True(d.DeleteLast(), "删除尾部元素")
	ast.True(d.InsertFront(4), "在首部添加 4")

	ast.Equal(4, d.GetFront(), "检查首部元素")
}
