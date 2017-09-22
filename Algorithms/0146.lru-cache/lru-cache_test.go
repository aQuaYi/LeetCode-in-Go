package Problem0146

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Constructor(t *testing.T) {
	ast := assert.New(t)

	cache := Constructor(2)

	cache.Put(1, 1)
	// [(1,1)]

	cache.Put(2, 2)
	// [(2,2), (1,1)]

	ast.Equal(1, cache.Get(1), "get 1 from [(2,2), (1,1)]")
	// [(1,1), (2,2)]

	cache.Put(3, 3)
	// [(3,3), (1,1)]

	ast.Equal(-1, cache.Get(2), "get 2 from [(3,3), (1,1)]")

	cache.Put(4, 4)
	// [(4,4), (3,3)]

	ast.Equal(-1, cache.Get(1), "get 1 from [(4,4), (3,3)]")

	ast.Equal(3, cache.Get(3), "get 3 from [(4,4), (3,3)]")
	// [(3,3), (4,4)]

	ast.Equal(4, cache.Get(4), "get 4 from [(3,3), (4,4)]")
	// [(4,4), (3,3)]
}
