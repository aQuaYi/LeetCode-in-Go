package problem0146

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// capacity is 3
func Test_3(t *testing.T) {
	ast := assert.New(t)

	cache := Constructor(3)

	cache.Put(0, 0)
	// [(0,0)]

	cache.Put(1, 1)
	// [(1,1), (0,0)]

	cache.Put(2, 2)
	// [(2,2), (1,1), (0,0)]

	ast.Equal(1, cache.Get(1), "get 1 from [(2,2), (1,1), (0,0)]")
	// [(1,1), (2,2), (0,0)]

	cache.Put(3, 3)
	// [(3,3), (1,1), (2,2)]

	ast.Equal(-1, cache.Get(0), "get 2 from [(3,3), (1,1), (2,2)]")

	cache.Put(4, 4)
	// [(4,4), (3,3), (1,1)]

	ast.Equal(-1, cache.Get(2), "get 1 from [(4,4), (3,3), (1,1)]")

	ast.Equal(3, cache.Get(3), "get 3 from [(4,4), (3,3), (1,1)]")
	// [(3,3), (4,4), (1,1)]

	ast.Equal(3, cache.Get(3), "get 3 from [(3,3), (4,4), (1,1)]")
	// [(3,3), (4,4), (1,1)]

	ast.Equal(4, cache.Get(4), "get 4 from [(3,3), (4,4), (1,1)]")
	// [(4,4), (3,3), (1,1)]

	ast.Equal(1, cache.Get(1), "get 4 from [(3,3), (4,4), (1,1)]")
	// [(1,1), (4,4), (3,3)]

	cache.Put(5, 5)
	// [(5,5), (1,1), (4,4)]
}

// capacity is 1
func Test_1(t *testing.T) {
	ast := assert.New(t)

	cache := Constructor(1)

	cache.Put(0, 0)
	// [(0,0)]

	cache.Put(1, 1)
	// [(1,1)]

	cache.Put(2, 2)
	// [(2,2)]

	ast.Equal(-1, cache.Get(1), "get 1 from [(2,2)]")
	// [(1,1), (2,2), (0,0)]

	cache.Put(3, 3)
	// [(3,3)]

	cache.Put(3, 33)
	// [(3,33)]
}
