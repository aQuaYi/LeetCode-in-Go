package problem0946

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	pushed []int
	popped []int
	ans    bool
}{

	{
		[]int{0, 2, 1},
		[]int{0, 1, 2},
		true,
	},

	{
		[]int{1, 2, 3, 4, 5},
		[]int{4, 5, 3, 2, 1},
		true,
	},

	{
		[]int{1, 2, 3, 4, 5},
		[]int{4, 3, 5, 1, 2},
		false,
	},

	// 可以有多个 testcase
}

func Test_validateStackSequences(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, validateStackSequences(tc.pushed, tc.popped), "输入:%v", tc)
	}
}

func Benchmark_validateStackSequences(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			validateStackSequences(tc.pushed, tc.popped)
		}
	}
}
