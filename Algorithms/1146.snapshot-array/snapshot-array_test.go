package problem1146

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SnapshotArray(t *testing.T) {
	ast := assert.New(t)

	sa := Constructor(3)

	sa.Set(0, 5)

	ast.Equal(0, sa.Snap())

	sa.Set(0, 6)

	ast.Equal(5, sa.Get(0, 0))

	sa.Set(1, 7)

	ast.Equal(0, sa.Get(1, 0))

	ast.Equal(0, sa.Get(2, 0))

}

func Test_SnapshotArray_2(t *testing.T) {
	ast := assert.New(t)

	sa := Constructor(3)

	sa.Set(1, 14)

	ast.Equal(0, sa.Snap())
	ast.Equal(1, sa.Snap())
	ast.Equal(2, sa.Snap())

	sa.Set(0, 14)

	ast.Equal(3, sa.Snap())
	ast.Equal(4, sa.Snap())

	sa.Set(2, 0)
	sa.Set(0, 5)

	ast.Equal(14, sa.Get(0, 3))

	ast.Equal(5, sa.Snap())
	ast.Equal(6, sa.Snap())

}
