package problem0676

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MagicDictionary(t *testing.T) {
	ast := assert.New(t)

	md := Constructor()

	dict := []string{"hello", "leetcode"}
	md.BuildDict(dict)

	ast.False(md.Search("hello"))

	ast.True(md.Search("hhllo"))

	ast.False(md.Search("hell"))

	ast.False(md.Search("leetcoded"))
}

func Test_MagicDictionary_2(t *testing.T) {
	ast := assert.New(t)

	md := Constructor()

	dict := []string{"hello", "hallo", "leetcode"}
	md.BuildDict(dict)

	ast.True(md.Search("hello"), "hello 可以从 hallo 变化过来")
}
