package problem0324

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
}{

	{[]int{1, 5, 1, 1, 6, 4}},
	{[]int{1, 3, 2, 2, 3, 1}},

	// 可以有多个 testcase
}

func check(nums []int) bool {
	for i := 0; i < len(nums); i += 2 {
		if 0 <= i-1 && nums[i-1] <= nums[i] {
			return false
		}
		if i+1 < len(nums) && nums[i] >= nums[i+1] {
			return false
		}
	}
	return true
}

func Test_wiggleSort(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		wiggleSort(tc.nums)
		ast.True(check(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_wiggleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			wiggleSort(tc.nums)
		}
	}
}
