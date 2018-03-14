package problem0605

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Problem0605(t *testing.T) {
	ast := assert.New(t)

	questions := []struct {
		p1       []int
		p2       int
		expected bool
		errMsg   string
	}{
		{
			[]int{1},
			0,
			true,
			"正例测试失败",
		},
		{
			[]int{0},
			3,
			false,
			"反例测试失败",
		},
		{
			[]int{0, 0},
			1,
			true,
			"正例测试失败",
		},
		{
			[]int{0, 0},
			3,
			false,
			"反例测试失败",
		},
		{
			[]int{0, 0, 0, 0},
			3,
			false,
			"反例测试失败",
		},
		{
			[]int{1, 1, 1, 1, 1, 0, 0},
			1,
			true,
			"正例测试失败",
		},
		{
			[]int{1, 0, 0, 0, 0, 0, 1},
			0,
			true,
			"正例测试失败",
		},
		{
			[]int{1, 0, 0, 0, 0, 0, 1},
			2,
			true,
			"正例测试失败",
		},
		{
			[]int{1, 0, 0, 0, 1},
			2,
			false,
			"反例测试失败",
		},
		{
			[]int{0, 0, 0, 0, 0},
			2,
			true,
			"正例测试失败",
		},
		{
			[]int{0},
			1,
			true,
			"正例测试失败",
		},
		{
			[]int{0, 0, 0, 1},
			1,
			true,
			"正例测试失败",
		},
		{
			[]int{1, 0, 0, 0, 1},
			1,
			true,
			"正例测试失败",
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range questions {
		ast.Equal(q.expected, canPlaceFlowers(q.p1, q.p2), q.errMsg)
	}
}
