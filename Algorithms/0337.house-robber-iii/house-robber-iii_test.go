package Problem0337

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	pre, in []int
	ans     int
}{

	{
		[]int{5, 2, 4, 1, 3, 1000},
		[]int{2, 4, 5, 1, 3, 1000},
		1009,
	},

	{
		[]int{5, 2, 4, 1, 3},
		[]int{2, 4, 5, 1, 3},
		12,
	},

	{
		[]int{3, 4, 1, 3, 5, 1},
		[]int{1, 4, 3, 3, 5, 1},
		9,
	},

	// 可以有多个 testcase
}

func Test_rob(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := kit.PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.ans, rob(root), "输入:%v", tc)
	}
}

func Benchmark_rob(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.PreIn2Tree(tc.pre, tc.in)
			rob(root)
		}
	}
}
