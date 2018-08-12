package problem0706

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_myFunc(t *testing.T) {
	ast := assert.New(t)

	m := Constructor()

	m.Put(1, 1)

	m.Put(2, 2)

	ast.Equal(1, m.Get(1))

	ast.Equal(-1, m.Get(3))

	m.Put(2, 1)

	ast.Equal(1, m.Get(2))

	m.Remove(2)

	ast.Equal(-1, m.Get(2))
}
