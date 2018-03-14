package problem0698

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	k    int
	ans  bool
}{

	{
		[]int{5, 4, 4, 4, 3},
		4,
		false,
	},

	{
		[]int{4, 1, 1, 1},
		4,
		false,
	},

	{
		[]int{4, 3, 2, 3, 5, 2, 1},
		4,
		true,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, canPartitionKSubsets(tc.nums, tc.k), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			canPartitionKSubsets(tc.nums, tc.k)
		}
	}
}
