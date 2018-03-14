package problem0225

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MyStack(t *testing.T) {
	ast := assert.New(t)

	s := Constructor()
	ast.True(s.Empty(), "检查新建的 s 是否为空")

	start, end := 0, 10

	for i := start; i < end; i++ {
		s.Push(i)
		ast.Equal(i, s.Top(), "查看 s.Top()")
	}

	for i := end - 1; i >= start; i-- {
		ast.Equal(i, s.Pop(), "从 s 中 pop 出数来。")
	}

	ast.True(s.Empty(), "检查 Pop 完毕后的 s 是否为空")
}

func Test_Queue(t *testing.T) {
	ast := assert.New(t)

	q := NewQueue()
	ast.True(q.IsEmpty(), "检查新建的 q 是否为空")

	start, end := 0, 100

	for i := start; i < end; i++ {
		q.Push(i)
		ast.Equal(i-start+1, q.Len(), "Push 后检查 q 的长度。")
	}

	for i := start; i < end; i++ {
		ast.Equal(i, q.Pop(), "从 q 中 pop 出数来。")
	}

	ast.True(q.IsEmpty(), "检查 Pop 完毕后的 q 是否为空")
}
