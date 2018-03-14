package problem0330

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	n    int
	ans  int
}{

	{[]int{1, 3}, 6, 1},

	{[]int{1, 5, 10}, 20, 2},

	{[]int{1, 2, 2}, 5, 0},

	{[]int{1, 2, 2, 2}, 7, 0},

	{[]int{1, 2, 2}, 7, 1},

	{[]int{1, 2, 31, 33}, 2147483647, 28},

	{[]int{1, 2, 14, 15, 17, 18}, 2147483647, 27},

	{[]int{1, 2, 14, 15, 17, 18}, 65, 2},


	{[]int{1, 2, 63	}, 65, 4},

	//
}

func Test_minPatches(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, minPatches(tc.nums, tc.n), "输入:%v", tc)
	}
}

func Benchmark_minPatches(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minPatches(tc.nums, tc.n)
		}
	}
}
