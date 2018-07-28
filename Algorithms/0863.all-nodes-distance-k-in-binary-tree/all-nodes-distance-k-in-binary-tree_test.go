package problem0863

import (
	"fmt"
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
		[]int{3, 5, 1, 6, 2, 0, 8, kit.NULL, kit.NULL, 7, 4},
		5,
		2,
		[]int{7, 4, 1},
	},

	// 可以有多个 testcase
}

func Test_distanceK(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		root := kit.Ints2TreeNode(tc.root)
		target := kit.GetTargetNode(root, tc.target)

		ast.Equal(tc.ans, distanceK(root, target, tc.K), "输入:%v", tc)
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
