package problem0707

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_myLinkedList(t *testing.T) {
	ast := assert.New(t)

	l := Constructor()

	l.AddAtHead(1)
	l.AddAtTail(3)

	l.AddAtIndex(5, 5)
	ast.Equal(-1, l.Get(5))

	l.AddAtIndex(1, 2)

	ast.Equal(2, l.Get(1))

	l.DeleteAtIndex(1)

	ast.Equal(3, l.Get(1))
}
func Test_myLinkedList_2(t *testing.T) {
	ast := assert.New(t)

	l := Constructor()

	l.AddAtTail(2)
	l.AddAtHead(1)
	l.AddAtTail(4)
	l.AddAtTail(5)

	l.AddAtIndex(2, 3)
	l.AddAtIndex(0, 0)
	l.AddAtIndex(6, 6)

	ast.Equal(1, l.Get(1))

	l.DeleteAtIndex(5)
	l.DeleteAtIndex(8)

	ast.Equal(6, l.Get(5))

}
func Test_myLinkedList_3(t *testing.T) {
	ast := assert.New(t)

	l := Constructor()

	ast.Equal(-1, l.Get(0))

	l.AddAtIndex(1, 2)

	ast.Equal(-1, l.Get(0))
	ast.Equal(-1, l.Get(1))

	l.AddAtIndex(0, 1)

	ast.Equal(1, l.Get(0))
	ast.Equal(-1, l.Get(1))

}
