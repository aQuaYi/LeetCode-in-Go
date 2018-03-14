package problem0155

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Problem155_1(t *testing.T) {
	ast := assert.New(t)

	s := Constructor()

	s.Push(-2)
	// [-2]
	s.Push(0)
	// [-2, 0]
	s.Push(-3)
	// [-2, 0, -3]
	ast.Equal(-3, s.GetMin(), "get min from [-2, 0, -3]")
	// [-2, 0, -3]
	s.Pop()
	// [-2, 0]
	ast.Equal(0, s.Top(), "get top from [-2, 0]")
	// [-2, 0]
	ast.Equal(-2, s.GetMin(), "get min from [-2, 0]")
	// [-2, 0]
}
func Test_Problem155_2(t *testing.T) {
	ast := assert.New(t)

	s := Constructor()

	s.Push(-2)
	// [-2]
	s.Push(0)
	// [-2, 0]
	s.Push(-1)
	// [-2, 0, -1]
	ast.Equal(-2, s.GetMin(), "get min from [-2, -1, 0]")
	// [-2, 0, -1]
	ast.Equal(-1, s.Top(), "get top from [-2, -1, 0]")
	// [-2, 0, -1]
	s.Pop()
	// [-2, 0]
	ast.Equal(-2, s.GetMin(), "get top from [0, -1]")
	// [-2, 0]
	s.Pop()
	// [-2]
}
