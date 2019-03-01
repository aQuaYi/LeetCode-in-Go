package problem0945

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	ans int
}{

	{
		[]int{},
		0,
	},

	{
		[]int{1, 2, 2},
		1,
	},

	{
		[]int{3, 2, 1, 2, 1, 7},
		6,
	},

	// 可以有多个 testcase
}

func Test_minIncrementForUnique(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, minIncrementForUnique(tc.A), "输入:%v", tc)
	}
}

func Benchmark_minIncrementForUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minIncrementForUnique(tc.A)
		}
	}
}
