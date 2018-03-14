package problem0697

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  int
}{

	{
		[]int{4},
		1,
	},

	{
		[]int{1, 2, 3, 4},
		1,
	},

	{
		[]int{1, 2, 2, 3, 1},
		2,
	},

	{
		[]int{1, 2, 2, 3, 1, 4, 2},
		6,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findShortestSubArray(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findShortestSubArray(tc.nums)
		}
	}
}
