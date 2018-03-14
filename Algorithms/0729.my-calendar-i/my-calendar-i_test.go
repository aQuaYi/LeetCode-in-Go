package problem0729

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MyCanlendar(t *testing.T) {
	ast := assert.New(t)

	mc := Constructor()

	ast.True(mc.Book(10, 20))

	ast.False(mc.Book(15, 25))

	ast.True(mc.Book(20, 30))
}
