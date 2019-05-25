package problem0981

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Example_1(t *testing.T) {
	ast := assert.New(t)

	kv := Constructor()
	kv.Set("foo", "bar", 1)

	ast.Equal("bar", kv.Get("foo", 1))
	ast.Equal("bar", kv.Get("foo", 3))

	kv.Set("foo", "bar2", 4)

	ast.Equal("bar2", kv.Get("foo", 4))
	ast.Equal("bar2", kv.Get("foo", 5))
}

func Test_Example_2(t *testing.T) {
	ast := assert.New(t)

	kv := Constructor()
	kv.Set("love", "high", 10)
	kv.Set("love", "low", 20)

	ast.Equal("", kv.Get("love", 5))
	ast.Equal("high", kv.Get("love", 10))
	ast.Equal("high", kv.Get("love", 15))
	ast.Equal("low", kv.Get("love", 20))
	ast.Equal("low", kv.Get("love", 25))

}
