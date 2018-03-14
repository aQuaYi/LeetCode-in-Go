package problem0376

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
		[]int{1},
		1,
	},

	{
		[]int{1, 2},
		2,
	},

	{
		[]int{1, 7, 4, 9, 2, 5},
		6,
	},

	{
		[]int{1, 1, 1, 1},
		1,
	},

	{
		[]int{1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2},
		2,
	},

	{
		[]int{7, 4, 9, 2, 5},
		5,
	},

	{
		[]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8},
		7,
	},

	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		2,
	},

	// 可以有多个 testcase
}

func Test_wiggleMaxLength(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, wiggleMaxLength(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_wiggleMaxLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			wiggleMaxLength(tc.nums)
		}
	}
}
