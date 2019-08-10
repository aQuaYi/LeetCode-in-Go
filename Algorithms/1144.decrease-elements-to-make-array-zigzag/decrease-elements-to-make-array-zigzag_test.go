package problem1144

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  int
}{

	{
		[]int{1, 2, 3},
		2,
	},

	{
		[]int{9, 6, 1, 6, 2},
		4,
	},

	// 可以有多个 testcase
}

func Test_movesToMakeZigzag(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, movesToMakeZigzag(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_movesToMakeZigzag(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			movesToMakeZigzag(tc.nums)
		}
	}
}
