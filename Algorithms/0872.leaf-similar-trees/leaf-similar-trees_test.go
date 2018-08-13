package problem0872

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root1 []int
	root2 []int
	ans   bool
}{

	{
		[]int{44, 79, 25, kit.NULL, kit.NULL, 112, 7, 74, 49, 2, 122},
		[]int{38, 86, 120, 49, 54, 2, 122, kit.NULL, kit.NULL, 74, 79},
		false,
	},

	{
		[]int{3, 5, 1, 6, 2, 9, 8, kit.NULL, kit.NULL, 7, 4},
		[]int{3, 5, 1, 6, 7, 4, 2, kit.NULL, kit.NULL, kit.NULL, kit.NULL, kit.NULL, kit.NULL, 9, 8},
		true,
	},

	{
		[]int{3, 5, 1, 6, 2, 9, 8, kit.NULL, kit.NULL, 7, 4},
		[]int{3, 5, 1, 6, 7, 4, 2, kit.NULL, kit.NULL, kit.NULL, kit.NULL, kit.NULL, kit.NULL, 9, 9},
		false,
	},

	{
		[]int{3, 5, 1, 6, 2, 9, 8, 3, kit.NULL, 7, 4},
		[]int{3, 5, 1, 6, 7, 4, 2, kit.NULL, kit.NULL, kit.NULL, kit.NULL, kit.NULL, kit.NULL, 9, 9},
		false,
	},

	// 可以有多个 testcase
}

func Test_leafSimilar(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root1 := kit.Ints2TreeNode(tc.root1)
		root2 := kit.Ints2TreeNode(tc.root2)
		ast.Equal(tc.ans, leafSimilar(root1, root2), "输入:%v", tc)
	}
}

func Benchmark_leafSimilar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root1 := kit.Ints2TreeNode(tc.root1)
			root2 := kit.Ints2TreeNode(tc.root2)
			leafSimilar(root1, root2)
		}
	}
}
