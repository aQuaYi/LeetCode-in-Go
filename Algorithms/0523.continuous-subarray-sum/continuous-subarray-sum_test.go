package problem0523

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
		[]int{1, 2},
		4,
		false,
	},

	{
		[]int{0, 0},
		0,
		true,
	},

	{
		[]int{23, 2, 4, 6, 7},
		6,
		true,
	},

	{
		[]int{1, 2, 3, 4},
		5,
		true,
	},

	{
		[]int{23, 2, 6, 4, 7},
		6,
		true,
	},

	{
		[]int{23, 2, 6, 4, 7},
		42,
		true,
	},

	// 可以有多个 testcase
}

func Test_checkSubarraySum(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, checkSubarraySum(tc.nums, tc.k), "输入:%v", tc)
	}
}

func Benchmark_checkSubarraySum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			checkSubarraySum(tc.nums, tc.k)
		}
	}
}
