package Problem0160

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	questions := map[string]string{
		"": "",
	}

	for expected, parameter := range questions {
		ast.Equal(expected, parameter, "输入:%!s(MISSING)", parameter)
	}
}
