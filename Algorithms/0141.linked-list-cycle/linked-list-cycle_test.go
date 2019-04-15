package problem0141

import (
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	head []int
	pos  int
	ans  bool
}{

	{
		[]int{3, 2, 0, -4},
		1,
		true,
	},

	{
		[]int{1, 2},
		0,
		true,
	},

	{
		[]int{},
		-1,
		false,
	},

	{
		[]int{1},
		-1,
		false,
	},

	// 可以有多个 testcase
}

func Test_hasCycle(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		head := kit.Ints2ListWithCycle(tc.head, tc.pos)
		ast.Equal(tc.ans, hasCycle(head), "输入:%v", tc)
	}
}

func Benchmark_hasCycle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			head := kit.Ints2ListWithCycle(tc.head, tc.pos)
			hasCycle(head)
		}
	}
}
