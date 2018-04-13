package problem0795

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	L   int
	R   int
	ans int
}{

	{
		[]int{2, 2, 4, 3},
		2,
		3,
		4,
	},

	{
		[]int{2, 1, 4, 3},
		2,
		3,
		3,
	},

	// 可以有多个 testcase
}

func Test_numSubarrayBoundedMax(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, numSubarrayBoundedMax(tc.A, tc.L, tc.R), "输入:%v", tc)
	}
}

func Benchmark_numSubarrayBoundedMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numSubarrayBoundedMax(tc.A, tc.L, tc.R)
		}
	}
}
