package problem0239

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	k    int
	ans  []int
}{

	{
		[]int{7, 2, 4},
		2,
		[]int{7, 4},
	},

	{
		[]int{},
		0,
		[]int{},
	},

	{
		[]int{2, 1, 3, 4, 6, 3, 8, 9, 10, 12, 56},
		4,
		[]int{4, 6, 6, 8, 9, 10, 12, 56},
	},

	{
		[]int{1, 3, -1, -3, 5, 3, 6, 7},
		1,
		[]int{1, 3, -1, -3, 5, 3, 6, 7},
	},

	{
		[]int{1, 3, -1, -3, 5, 3, 6, 7},
		3,
		[]int{3, 3, 5, 5, 6, 7},
	},

	// 可以有多个 testcase
}

func Test_maxSlidingWindow(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, maxSlidingWindow(tc.nums, tc.k), "输入:%v", tc)
	}
}

func Benchmark_maxSlidingWindow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxSlidingWindow(tc.nums, tc.k)
		}
	}
}
