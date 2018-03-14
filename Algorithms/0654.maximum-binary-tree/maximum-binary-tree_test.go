package problem0654

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	post []int
}{

	{
		[]int{3, 2, 1, 6, 0, 5},
		[]int{1, 2, 3, 0, 5, 6},
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.post, kit.Tree2Postorder(constructMaximumBinaryTree(tc.nums)), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			constructMaximumBinaryTree(tc.nums)
		}
	}
}
