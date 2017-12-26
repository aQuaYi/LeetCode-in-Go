package Problem0713

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	k    int
	ans  int
}{

	{
		[]int{1, 2, 3},
		0,
		0,
	},

	{
		[]int{10, 5, 2, 6},
		100,
		8,
	},

	// 可以有多个 testcase
}

func Test_numSubarrayProductLessThanK(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, numSubarrayProductLessThanK(tc.nums, tc.k), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numSubarrayProductLessThanK(tc.nums, tc.k)
		}
	}
}
