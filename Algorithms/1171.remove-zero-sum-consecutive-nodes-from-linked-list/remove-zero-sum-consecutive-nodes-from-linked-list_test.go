package problem1171

import (
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	head []int
	ans  []int
}{

	{
		[]int{0, 1, 1, -3, 6, -3, 0, 1},
		[]int{1, 1, 1},
	},

	{
		[]int{0, 1, 1, -3, 6, -3, 1},
		[]int{1, 1, 1},
	},

	{
		[]int{1, 1, -3, 6, -3, 1},
		[]int{1, 1, 1},
	},

	{
		[]int{1, 2, -2, 5, -3, 1},
		[]int{1, 5, -3, 1},
	},

	{
		[]int{1, 2, -8, 3, 1},
		[]int{1, 2, -8, 3, 1},
	},

	{
		[]int{1, 2, -3, 3, 1},
		[]int{3, 1},
	},

	{
		[]int{1, 2, 3, -3, 4},
		[]int{1, 2, 4},
	},

	{
		[]int{1, 2, 3, -3, -2},
		[]int{1},
	},

	// 可以有多个 testcase
}

func Test_removeZeroSumSublists(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		head := kit.Ints2List(tc.head)
		ans := kit.List2Ints(removeZeroSumSublists(head))
		a.Equal(tc.ans, ans, "输入:%v", tc)
	}
}

func Benchmark_removeZeroSumSublists(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			head := kit.Ints2List(tc.head)
			removeZeroSumSublists(head)
		}
	}
}
