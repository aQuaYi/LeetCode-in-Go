package problem0895

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FreqStack_1(t *testing.T) {
	ast := assert.New(t)
	fs := Constructor()
	pushes := []int{5, 7, 5, 7, 4, 5}
	for _, p := range pushes {
		fs.Push(p)
	}
	expecteds := []int{5, 7, 5, 4}
	for _, expected := range expecteds {
		actual := fs.Pop()
		ast.Equal(expected, actual)
	}
}

func Test_FreqStack_2(t *testing.T) {
	ast := assert.New(t)
	fs := Constructor()
	//
	pushes := []int{1, 1, 1, 2}
	for _, p := range pushes {
		fs.Push(p)
	}
	expecteds := []int{1, 1}
	for _, expected := range expecteds {
		actual := fs.Pop()
		ast.Equal(expected, actual)
	}
	//
	pushes = []int{2, 2, 1}
	for _, p := range pushes {
		fs.Push(p)
	}
	expecteds = []int{2, 1, 2}
	for _, expected := range expecteds {
		actual := fs.Pop()
		ast.Equal(expected, actual)
	}
}
