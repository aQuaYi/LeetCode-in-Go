package problem0653

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	pre, in []int
	k       int
	ans     bool
}{

	{
		[]int{2, 1, 3},
		[]int{1, 2, 3},
		4,
		true,
	},

	{
		[]int{1},
		[]int{1},
		2,
		false,
	},

	{
		[]int{5, 3, 2, 4, 6, 7},
		[]int{2, 3, 4, 5, 6, 7},
		28,
		false,
	},

	{
		[]int{5, 3, 2, 4, 6, 7},
		[]int{2, 3, 4, 5, 6, 7},
		9,
		true,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := kit.PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.ans, findTarget(root, tc.k), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.PreIn2Tree(tc.pre, tc.in)
			findTarget(root, tc.k)
		}
	}
}
