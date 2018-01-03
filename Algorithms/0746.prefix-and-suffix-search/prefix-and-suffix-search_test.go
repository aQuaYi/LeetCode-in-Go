package Problem0746

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_wordFilter(t *testing.T) {
	ast := assert.New(t)

	words := []string{"apple", "banana", "abbbe"}
	w := Constructor(words)

	ast.Equal(0, w.F("ap", "e"))

	ast.Equal(1, w.F("ba", ""))

	ast.Equal(2, w.F("a", "e"))

	ast.Equal(2, w.F("", ""))

}
