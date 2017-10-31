package kit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NestedInteger(t *testing.T) {
	ast := assert.New(t)

	n := NestedInteger{}

	ast.True(n.IsInteger())

	n.SetInteger(1)
	ast.Equal(1, n.GetInteger())

	elem := NestedInteger{num: 1}

	expected := NestedInteger{
		num: 1,
		ns:  []*NestedInteger{&elem},
	}
	n.Add(elem)

	ast.Equal(expected, n)

	ast.Equal(expected.ns, n.GetList())
}
