package problem0457

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  bool
}{

	{
		[]int{-1, 2},
		false,
	},

	{
		[]int{2, -1, 1, -2, -2},
		false,
	},

	{
		[]int{-2, 1, -1, -2, -2},
		false,
	},

	{
		[]int{2, -1, 1, 2, 2},
		true,
	},

	{
		[]int{-1, -1},
		true,
	},

	// 可以有多个 testcase
}

func Test_find132pattern(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, circularArrayLoop(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_find132pattern(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			circularArrayLoop(tc.nums)
		}
	}
}
