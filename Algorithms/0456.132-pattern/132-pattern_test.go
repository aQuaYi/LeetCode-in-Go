package Problem0456

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
