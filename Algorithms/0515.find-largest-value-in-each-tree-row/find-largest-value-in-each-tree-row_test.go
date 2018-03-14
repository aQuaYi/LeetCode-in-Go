package problem0515

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	pre, in []int
	ans     []int
}{

	{
		[]int{},
		[]int{},
		[]int{},
	},

	{
		[]int{1, 3, 5, 6, 2, 9},
		[]int{5, 3, 6, 1, 2, 9},
		[]int{1, 3, 9},
	},

	// 可以有多个 testcase
}

func Test_largestValues(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := kit.PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.ans, largestValues(root), "输入:%v", tc)
	}
}

func Benchmark_largestValues(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.PreIn2Tree(tc.pre, tc.in)
			largestValues(root)
		}
	}
}
