package problem0456

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
		[]int{9, 8, 7, 6, 5, 4, 3, 2, 1},
		false,
	},

	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
		false,
	},

	{
		[]int{9, 11, 8, 9, 10, 7, 9},
		true,
	},

	{
		[]int{1, 0, 1, -4, -3},
		false,
	},

	{
		[]int{3, 4},
		false,
	},

	{
		[]int{1, 2, 3, 4},
		false,
	},

	{
		[]int{3, 1, 4, 2},
		true,
	},

	{
		[]int{-1, 3, 2, 0},
		true,
	},

	// 可以有多个 testcase
}

func Test_find132pattern(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, find132pattern(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_find132pattern(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			find132pattern(tc.nums)
		}
	}
}
