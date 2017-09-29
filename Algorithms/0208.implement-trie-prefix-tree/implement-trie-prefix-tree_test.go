package Problem0208

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Constructor(t *testing.T) {
	ast := assert.New(t)

	trie := Constructor()

	trie.Insert("abc")
	// ["abc"]

	ast.True(trie.Search("abc"), "Search abc in [abc]")
	// ["abc"]

	ast.False(trie.Search("abcd"), "Search abcd in [abc]")
	// ["abc"]

	ast.True(trie.StartsWith("abc"), "Search startWith abc in [abc]")
	// ["abc"]

	ast.False(trie.StartsWith("bc"), "Search startWith bc in [abc]")
	// ["abc"]
}
