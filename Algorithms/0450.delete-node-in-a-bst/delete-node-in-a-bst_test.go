package Problem0450

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	pre, in []int
	key     int
	ans     []int
}{

	{
		[]int{1, 2},
		[]int{1, 2},
		1,
		[]int{2},
	},

	{
		[]int{0},
		[]int{0},
		0,
		nil,
	},

	{
		[]int{5, 3, 2, 4, 6, 7},
		[]int{2, 3, 4, 5, 6, 7},
		3,
		[]int{5, 2, 4, 6, 7},
	},

	// 可以有多个 testcase
}

func Test_deleteNode(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := kit.PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.ans, kit.Tree2Preorder(deleteNode(root, tc.key)), "输入:%v", tc)
	}
}

func Benchmark_deleteNode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.PreIn2Tree(tc.pre, tc.in)
			deleteNode(root, tc.key)
		}
	}
}
