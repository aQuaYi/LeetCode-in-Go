package problem0142

import (
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	ints []int
	pos  int
}{

	{
		[]int{1},
		-1,
	},

	{
		[]int{1, 2, 3},
		-1,
	},

	{
		[]int{3, 2, 0, -4},
		1,
	},

	{
		[]int{1, 2},
		0,
	},

	// 可以有多个 testcase
}

func Test_detectCycle(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		head := kit.Ints2ListWithCycle(tc.ints, tc.pos)
		var ans *ListNode
		if tc.pos >= 0 {
			ans = head.GetNodeWith(tc.ints[tc.pos])
		}
		ast.Equal(ans, detectCycle(head), "输入:%v", tc)
	}
}

func Benchmark_detectCycle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			head := kit.Ints2ListWithCycle(tc.ints, tc.pos)
			detectCycle(head)
		}
	}
}
