package Problem0239

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
