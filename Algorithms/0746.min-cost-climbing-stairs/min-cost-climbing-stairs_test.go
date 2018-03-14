package problem0746

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	cost []int
	ans  int
}{

	{
		[]int{10, 15, 20},
		15,
	},

	{
		[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
		6,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, minCostClimbingStairs(tc.cost), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minCostClimbingStairs(tc.cost)
		}
	}
}
