package Problem0001

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
	two int
}

func (p para) String() string {
	res := fmt.Sprintf("%s", p.one)
	return res
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	expected := []int{1, 2}
	p := para{[]int{3, 2, 4}, 6}
	ast.Equal(expected, p.one, "输入:%!s(MISSING)", p)
}
