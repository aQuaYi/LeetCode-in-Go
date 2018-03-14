package problem0506

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  []string
}{

	{
		[]int{0, 4, 3, 2, 1},
		[]string{"5", "Gold Medal", "Silver Medal", "Bronze Medal", "4"},
	},

	{
		[]int{5, 4, 3, 2, 1},
		[]string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"},
	},

	// 可以有多个 testcase
}

func Test_findRelativeRanks(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findRelativeRanks(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_findRelativeRanks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findRelativeRanks(tc.nums)
		}
	}
}
