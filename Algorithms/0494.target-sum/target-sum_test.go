package problem0494

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	S    int
	ans  int
}{

	{
		[]int{1, 1, 1, 1, 1},
		8,
		0,
	},

	{
		[]int{1, 1, 1, 1, 1},
		4,
		0,
	},

	{
		[]int{1, 1, 1, 1, 1},
		3,
		5,
	},

	// 可以有多个 testcase
}

func Test_findTargetSumWays(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findTargetSumWays(tc.nums, tc.S), "输入:%v", tc)
	}
}

func Benchmark_findTargetSumWays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findTargetSumWays(tc.nums, tc.S)
		}
	}
}
