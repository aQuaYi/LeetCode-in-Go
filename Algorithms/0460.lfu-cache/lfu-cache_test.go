package problem0460

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LFUCache(t *testing.T) {
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

func Test_LFUCache2(t *testing.T) {
	ast := assert.New(t)

	cache := Constructor(0)

	cache.Put(0, 0)

	ast.Equal(-1, cache.Get(0), "没能正确处理好， cap <= 0 的情况")
}

func Test_LFUCache3(t *testing.T) {
	ast := assert.New(t)

	cache := Constructor(2)

	cache.Put(3, 1)
	cache.Put(2, 1)
	cache.Put(2, 2)
	cache.Put(4, 4)

	ast.Equal(2, cache.Get(2), "更新 2 的值后，其 frequence 没有增加")
}
