package problem0863

import (
	"fmt"
	"sort"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root   []int
	target int
	K      int
	ans    []int
}{

	{
		[]int{0, kit.NULL, 1, 2, 5, kit.NULL, 3, kit.NULL, kit.NULL, kit.NULL, 4},
		2,
		2,
		[]int{4, 5, 0},
	},

	{
		[]int{3, 5, 1, 6, 2, 0, 8, kit.NULL, kit.NULL, 7, 4},
		4,
		2,
		[]int{5, 7},
	},

	{
		[]int{3, 5, 1, 6, 2, 0, 8, kit.NULL, kit.NULL, 7, 4},
		5,
		1,
		[]int{2, 3, 6},
	},

	{
		[]int{0, kit.NULL, 1, kit.NULL, 2, kit.NULL, 3},
		1,
		2,
		[]int{3},
	},

	{
		[]int{3, 5, 1, 6, 2, 0, 8, kit.NULL, kit.NULL, 7, 4},
		5,
		0,
		[]int{5},
	},

	{
		[]int{3, 5, 1, 6, 2, 0, 8, kit.NULL, kit.NULL, 7, 4},
		5,
		2,
		[]int{7, 4, 1},
	},

	{
		[]int{3},
		3,
		0,
		[]int{3},
	},

	// 可以有多个 testcase
}

func Test_distanceK(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		root := kit.Ints2TreeNode(tc.root)
		target := kit.GetTargetNode(root, tc.target)
		ans := distanceK(root, target, tc.K)
		sort.Ints(ans)
		sort.Ints(tc.ans)
		ast.Equal(tc.ans, ans, "输入:%v", tc)
	}
}

func Benchmark_distanceK(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {

			root := kit.Ints2TreeNode(tc.root)
			target := kit.GetTargetNode(root, tc.target)

			distanceK(root, target, tc.K)
		}
	}
}
