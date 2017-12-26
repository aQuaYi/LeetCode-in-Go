package Problem0715

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	rm := Constructor()

	rm.AddRange(10, 20)

	rm.RemoveRange(14, 16)

	ast.True(rm.QueryRange(10, 14))

	ast.False(rm.QueryRange(13, 15))

	ast.True(rm.QueryRange(16, 17))
}
