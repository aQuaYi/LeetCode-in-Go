package Problem0211

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Constructor(t *testing.T) {
	ast := assert.New(t)

	d := Constructor()
	d.AddWord("bad")
	// [bad]
	d.AddWord("dad")
	// [bad, dad]
	d.AddWord("mad")
	// [bad, dad, mad]

	ast.False(d.Search("pad"), "search pad from [bad, dad, mad]")

	ast.True(d.Search("bad"), "search bad from [bad, dad, mad]")

	ast.True(d.Search(".ad"), "search .ad from [bad, dad, mad]")

	ast.True(d.Search("b.."), "search b.. from [bad, dad, mad]")
}
