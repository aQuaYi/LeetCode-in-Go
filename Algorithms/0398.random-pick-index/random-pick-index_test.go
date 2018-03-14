package problem0398

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Constructor(t *testing.T) {
	ast := assert.New(t)

	nums := []int{1, 2, 2, 3, 3, 3}
	// all index of 1
	is1 := []int{0}
	// all index of 2
	is2 := []int{1, 2}
	// all index of 3
	is3 := []int{3, 4, 5}

	s := Constructor(nums)

	ast.Contains(is1, s.Pick(1), "无法正确第返回 1 的索引号")

	ast.Contains(is2, s.Pick(2), "无法正确第返回 2 的索引号")

	ast.Contains(is3, s.Pick(3), "无法正确第返回 3 的索引号")

}
