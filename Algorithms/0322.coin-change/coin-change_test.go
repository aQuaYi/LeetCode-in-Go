package problem0322

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	coins  []int
	amount int
	ans    int
}{

	{[]int{1}, 1, 1},
	{[]int{1}, 0, 0},
	{[]int{1, 3, 5}, 8, 2},
	{[]int{1, 2}, 3, 2},
	{[]int{1, 2, 5}, 11, 3},
	{[]int{2}, 3, -1},

	// 可以有多个 testcase
}

func Test_coinChange(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, coinChange(tc.coins, tc.amount), "输入:%v", tc)
	}
}

func Benchmark_coinChange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			coinChange(tc.coins, tc.amount)
		}
	}
}
