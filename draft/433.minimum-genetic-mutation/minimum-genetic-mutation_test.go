package Problem433

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ok(t *testing.T) {
	ast := assert.New(t)

	expected := ""
	actual := ""

	ast.Equal(expected, actual, "与预期不符")
}
