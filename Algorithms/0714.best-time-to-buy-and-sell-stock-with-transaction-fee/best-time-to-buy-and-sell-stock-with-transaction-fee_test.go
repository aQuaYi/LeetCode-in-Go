package problem0714

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	prices []int
	fee    int
	ans    int
}{

	{
		[]int{1, 3, 7, 5, 10, 3},
		3,
		6,
	},

	{
		[]int{1, 3, 2, 8, 4, 9},
		2,
		8,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, maxProfit(tc.prices, tc.fee), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxProfit(tc.prices, tc.fee)
		}
	}
}
