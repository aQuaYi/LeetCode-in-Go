package problem0622

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_myFunc(t *testing.T) {
	ast := assert.New(t)

	m := Constructor(3)

	ast.True(m.IsEmpty(), "检查空数组的 IsEmpty()")
	ast.False(m.IsFull(), "检查空数组的 IsFull()")

	ast.False(m.DeQueue(), "从空数组中删除元素")

	ast.Equal(-1, m.Front(), "检查空数组的首部元素")
	ast.Equal(-1, m.Rear(), "检查空数组的尾部元素")

	for i := 1; i <= 3; i++ {
		ast.True(m.EnQueue(i), "往数组中添加 %d", i)
	}

	ast.False(m.EnQueue(4), "往满数组添加元素")

	ast.False(m.IsEmpty(), "检查满数组的 IsEmpty()")
	ast.True(m.IsFull(), "检查满数组的 IsFull()")

	ast.Equal(3, m.Rear(), "检查数组尾部的元素")

	ast.True(m.DeQueue(), "从非空数组中删除元素")

	ast.Equal([]int{2, 3}, m.queue, "DeQueue 后，检查数组中的元素")

	ast.True(m.EnQueue(4), "往数组中添加新的元素")

	ast.Equal([]int{2, 3, 4}, m.queue, "EnQueue 后，检查数组中的元素")

	ast.Equal(4, m.Rear(), "检查数组尾部的元素")

	ast.Equal(2, m.Front(), "检查数组首部的元素")
}
