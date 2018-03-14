package problem0450

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
		[]int{2, 0, 1, 33, 25, 11, 10, 4, 3, 9, 18, 12, 14, 24, 22, 31, 29, 26, 27, 30, 32, 40, 34, 36, 35, 39, 45, 43, 42, 44, 46, 48},
		[]int{0, 1, 2, 3, 4, 9, 10, 11, 12, 14, 18, 22, 24, 25, 26, 27, 29, 30, 31, 32, 33, 34, 35, 36, 39, 40, 42, 43, 44, 45, 46, 48},
		33,
		[]int{0, 1, 2, 3, 4, 9, 10, 11, 12, 14, 18, 22, 24, 25, 26, 27, 29, 30, 31, 32, 34, 35, 36, 39, 40, 42, 43, 44, 45, 46, 48},
	},

	{
		[]int{1, 2},
		[]int{2, 1},
		1,
		[]int{2},
	},

	{
		[]int{1, 2},
		[]int{1, 2},
		1,
		[]int{2},
	},

	{
		[]int{},
		[]int{},
		0,
		nil,
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
		[]int{2, 4, 5, 6, 7},
	},

	// 可以有多个 testcase
}

func Test_deleteNode(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := kit.PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.ans, kit.Tree2Inorder(deleteNode(root, tc.key)), "输入:%v", tc)
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
