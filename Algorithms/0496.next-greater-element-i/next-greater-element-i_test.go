package problem0496

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	findNums []int
	nums     []int
	ans      []int
}{

	{
		[]int{4, 1, 2},
		[]int{1, 3, 4, 2},
		[]int{-1, 3, -1},
	},

	{
		[]int{2, 4},
		[]int{1, 2, 3, 4},
		[]int{3, -1},
	},

	// 可以有多个 testcase
}

func Test_nextGreaterElement(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, nextGreaterElement(tc.findNums, tc.nums), "输入:%v", tc)
	}
}

func Benchmark_nextGreaterElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			nextGreaterElement(tc.findNums, tc.nums)
		}
	}
}
