package problem0606

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	pre, in []int
	ans     string
}{

	{
		[]int{1, 2, 4, 3},
		[]int{4, 2, 1, 3},
		"1(2(4))(3)",
	},

	{
		[]int{1, 2, 4, 3},
		[]int{2, 4, 1, 3},
		"1(2()(4))(3)",
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := kit.PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.ans, tree2str(root), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.PreIn2Tree(tc.pre, tc.in)
			tree2str(root)
		}
	}
}
