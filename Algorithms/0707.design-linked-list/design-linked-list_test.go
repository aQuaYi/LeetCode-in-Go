package problem0707

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_myLinkedList(t *testing.T) {
	ast := assert.New(t)

	l := Constructor()

	l.AddAtHead(1)
	l.AddAtHead(3)

	l.AddAtIndex(5, 5)
	ast.Equal(-1, l.Get(5))

	l.AddAtIndex(1, 2)

	ast.Equal(2, l.Get(1))

	l.DeleteAtIndex(1)

	ast.Equal(3, l.Get(1))
}
