package problem0669

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	pre, in []int
	L, R    int
	ansPost []int
}{

	{
		[]int{1, 0, 2},
		[]int{0, 1, 2},
		1,
		2,
		[]int{2, 1},
	},

	{
		[]int{3, 0, 2, 1, 4},
		[]int{0, 1, 2, 3, 4},
		1,
		3,
		[]int{1, 2, 3},
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := kit.PreIn2Tree(tc.pre, tc.in)
		ansPost := kit.Tree2Postorder(trimBST(root, tc.L, tc.R))
		ast.Equal(tc.ansPost, ansPost, "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.PreIn2Tree(tc.pre, tc.in)
			trimBST(root, tc.L, tc.R)
		}
	}
}
