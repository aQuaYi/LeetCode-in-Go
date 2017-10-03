package Problem0215

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
		[]int{3, 2, 1, 5, 6, 4},
		2,
		5,
	},

	// 可以有多个 testcase
}

func Test_findKthLargest(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findKthLargest(tc.nums, tc.k), "输入:%v", tc)
	}
}

func Benchmark_findKthLargest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findKthLargest(tc.nums, tc.k)
		}
	}
}
