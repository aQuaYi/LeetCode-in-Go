package problem0416

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  bool
}{
	{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 100}, false},

	{[]int{1, 2, 5}, false},

	{[]int{1, 5, 11, 5}, true},

	{[]int{1, 3, 5, 7, 9, 11}, true},

	{[]int{1, 2, 3, 5}, false},

	// 可以有多个 testcase
}

func Test_canPartition(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, canPartition(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_canPartition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			canPartition(tc.nums)
		}
	}
}
