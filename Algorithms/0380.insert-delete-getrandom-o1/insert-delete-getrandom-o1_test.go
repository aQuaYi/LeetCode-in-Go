package problem0380

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Problem0380(t *testing.T) {
	ast := assert.New(t)

	rs := Constructor()

	ast.True(rs.Insert(1), "插入1")

	ast.False(rs.Remove(2), "删除2")

	ast.True(rs.Insert(2), "插入2")

	ast.Contains([]int{1, 2}, rs.GetRandom(), "返回1或2")

	ast.True(rs.Remove(1), "删除1")

	ast.False(rs.Insert(2), "再一次插入2")

	ast.Equal(2, rs.GetRandom(), "从只有2的集合中随机取出2")

	length := 100
	result := make([]int, length)
	for i := 0; i < 100; i++ {
		rs.Insert(i)
		result[i] = i
	}
	ast.Contains(result, rs.GetRandom(), "随机获取0~100之间的数")
}
