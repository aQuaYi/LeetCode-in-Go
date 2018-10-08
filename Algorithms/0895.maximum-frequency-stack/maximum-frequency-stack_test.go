package problem0895

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FreqStack(t *testing.T) {
	ast := assert.New(t)

	fs := Constructor()

	pushes := []int{5, 7, 5, 7, 4, 5}
	for _, p := range pushes {
		fs.Push(p)
	}

	expecteds := []int{5, 7, 4, 5}
	for _, expected := range expecteds {
		actual := fs.Pop()
		ast.Equal(expected, actual)
	}

}
