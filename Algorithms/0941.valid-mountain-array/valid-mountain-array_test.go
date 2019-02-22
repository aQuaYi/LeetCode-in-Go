package problem0941

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	ans bool
}{

	{
		[]int{2, 1},
		false,
	},

	{
		[]int{3, 5, 5},
		false,
	},

	{
		[]int{0, 3, 2, 1},
		true,
	},

	// 可以有多个 testcase
}

func Test_validMountainArray(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, validMountainArray(tc.A), "输入:%v", tc)
	}
}

func Benchmark_validMountainArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			validMountainArray(tc.A)
		}
	}
}
