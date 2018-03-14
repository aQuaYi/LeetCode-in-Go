package problem0594

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

	{
		[]int{1, 1, 1, 1},
		0,
	},

	{
		[]int{1, 3, 2, 2, 5, 2, 3, 7},
		5,
	},

	// 可以有多个 testcase
}

func Test_findLHS(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findLHS(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_findLHS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findLHS(tc.nums)
		}
	}
}
