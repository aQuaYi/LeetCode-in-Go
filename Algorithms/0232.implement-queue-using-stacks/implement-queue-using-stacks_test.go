package problem0232

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MyQueue(t *testing.T) {
	ast := assert.New(t)

	q := Constructor()
	ast.True(q.Empty(), "检查新建的 q 是否为空")

	start, end := 0, 3

	for i := start; i < end; i++ {
		q.Push(i)
		ast.Equal(start, q.Peek(), "查看 q.Peek()")
	}

	for i := start; i < end; i++ {
		ast.Equal(i, q.Pop(), "从 q 中 pop 出数来。")
	}
}
