package Problem0307

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SumRange(t *testing.T) {
	ast := assert.New(t)

	nums := []int{1, 3, 5}

	na := Constructor(nums)

	ast.Equal(9, na.SumRange(0, 2), "update 前，SumRange(0, 2)")

	na.Update(1, 2)

	ast.Equal(8, na.SumRange(0, 2), "update 后，SumRange(0, 2)")
}

func Test_SumRange_2(t *testing.T) {
	ast := assert.New(t)

	nums := []int{-1}

	na := Constructor(nums)

	ast.Equal(-1, na.SumRange(0, 0), "update 前，SumRange(0, 0)")

	na.Update(0, 1)

	ast.Equal(1, na.SumRange(0, 0), "update 后，SumRange(0, 2)")
}
