package problem0312

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  int
}{
	{[]int{}, 0},
	{[]int{3, 1, 5, 8}, 167},
	{[]int{3}, 3},

	// 可以有多个 testcase
}

func Test_maxCoins(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, maxCoins(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_maxCoins(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxCoins(tc.nums)
		}
	}
}
