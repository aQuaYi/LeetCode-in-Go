package Problem0731

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MyCanlendar(t *testing.T) {
	ast := assert.New(t)

	mc := Constructor()

	ast.True(mc.Book(10, 20))
	ast.True(mc.Book(50, 60))
	ast.True(mc.Book(10, 40))

	ast.False(mc.Book(5, 15))

	ast.True(mc.Book(5, 10))
	ast.True(mc.Book(25, 55))
}
