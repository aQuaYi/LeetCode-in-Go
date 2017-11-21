package Problem0503

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  []int
}{

	{
		[]int{1, 1, 1, 1, 1},
		[]int{-1, -1, -1, -1, -1},
	},

	{
		[]int{4, 3, 2, 1},
		[]int{-1, 4, 4, 4},
	},

	{
		[]int{1, 2, 1},
		[]int{2, -1, 2},
	},

	// 可以有多个 testcase
}

func Test_nextGreaterElements(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, nextGreaterElements(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_nextGreaterElements(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			nextGreaterElements(tc.nums)
		}
	}
}
