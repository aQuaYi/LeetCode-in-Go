package problem0473

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  bool
}{

	{
		[]int{2, 2, 2, 3, 4, 4, 4, 5, 6},
		true,
	},

	{
		[]int{10, 6, 5, 5, 5, 3, 3, 3, 2, 2, 2, 2},
		true,
	},

	{
		[]int{2, 4, 4, 4},
		false,
	},

	{
		[]int{1, 1, 2, 2, 2},
		true,
	},

	{
		[]int{},
		false,
	},

	{
		[]int{3, 3, 3, 5, 4},
		false,
	},

	{
		[]int{3, 3, 3, 3, 4},
		false,
	},

	// 可以有多个 testcase
}

func Test_makesquare(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, makesquare(tc.nums), "输入:%v", tc)
	}
}

// 由于 analyze 提前了处理了一些情况
// 导致 dfs 中的一处 return false 无法覆盖
// 只好添加以下单元测试覆盖掉
// 这样是不对滴
func Test_dfs(t *testing.T) {
	ast := assert.New(t)

	nums := []int{1, 1, 1, 1}
	edges := []int{2, 2, 2, 2}

	actual := dfs(nums, edges, 4, 1)
	expected := false
	ast.Equal(expected, actual)

}

func Benchmark_makesquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			makesquare(tc.nums)
		}
	}
}
