package problem0309

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	prices []int
	ans    int
}{

	{[]int{1, 2, 3, 0, 2}, 3},
	{[]int{}, 0},
	{[]int{3, 3}, 0},

	// 可以有多个 testcase
}

func Test_maxProfit(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, maxProfit(tc.prices), "输入:%v", tc)
	}
}

func Benchmark_maxProfit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxProfit(tc.prices)
		}
	}
}
