package problem0129

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
		[]int{1, 2, 3, 4},
		[]int{2, 1, 3, 4},
		146,
	},

	{
		[]int{1, 2, 3},
		[]int{2, 1, 3},
		25,
	},

	// 可以有多个 testcase
}

func Test_sumNumbers(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		root := kit.PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.ans, sumNumbers(root), "输入:%v", tc)
	}
}

func Benchmark_sumNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			b.StopTimer()
			root := kit.PreIn2Tree(tc.pre, tc.in)
			b.StartTimer()
			sumNumbers(root)
		}
	}
}
