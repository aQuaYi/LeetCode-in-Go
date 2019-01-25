package problem0211

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Constructor(t *testing.T) {
	ast := assert.New(t)

	d := Constructor()

	ast.False(d.Search("."), "search . from []")

	d.AddWord("bad")
	// [bad]
	d.AddWord("dad")
	// [bad, dad]
	d.AddWord("ma")
	// [bad, dad, ma]

	ast.False(d.Search("pad"), "search pad from [bad, dad, ma]")

	ast.True(d.Search("bad"), "search bad from [bad, dad, ma]")

	ast.True(d.Search(".ad"), "search .ad from [bad, dad, ma]")

	ast.True(d.Search("b.."), "search b.. from [bad, dad, ma]")

	ast.False(d.Search("."), "search . from [bad, dad, ma]")

	ast.False(d.Search("b"), "search b from [bad, dad, ma]")
}
