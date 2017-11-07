package Problem0460

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Constructor(t *testing.T) {
	ast := assert.New(t)

	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	ast.Equal(1, cache.Get(1))

	cache.Put(3, 3)

	ast.Equal(-1, cache.Get(2))

	ast.Equal(3, cache.Get(3))

	cache.Put(4, 4)

	ast.Equal(-1, cache.Get(1))

	ast.Equal(3, cache.Get(3))

	ast.Equal(4, cache.Get(4))

}
