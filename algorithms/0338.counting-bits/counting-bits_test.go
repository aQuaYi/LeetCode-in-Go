package Problem0338

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
}

func (p para) String() string {
	res := fmt.Sprintf("%v", p.one)
	return res
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	questions := map[string]para{
		"": para{""},
	}

	for expected, p := range questions {
		ast.Equal(expected, p.one, "输入:%s", p)
	}
}
