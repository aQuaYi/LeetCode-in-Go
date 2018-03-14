package problem0689

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
		[]int{17, 7, 19, 11, 1, 19, 17, 6, 13, 18, 2, 7, 12, 16, 16, 18, 9, 3, 19, 5},
		6,
		[]int{0, 6, 13},
	},

	{
		[]int{1, 2, 1, 2, 6, 7, 5, 1},
		2,
		[]int{0, 3, 5},
	},

	// 可以有多个 testcase
}

func Test_maxSumOfThreeSubarrays(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, maxSumOfThreeSubarrays(tc.nums, tc.k), "输入:%v", tc)
	}
}

func Benchmark_maxSumOfThreeSubarrays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxSumOfThreeSubarrays(tc.nums, tc.k)
		}
	}
}
