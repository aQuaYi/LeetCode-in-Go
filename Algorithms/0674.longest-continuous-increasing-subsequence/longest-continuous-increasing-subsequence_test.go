package problem0674

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
		[]int{},
		0,
	},

	{
		[]int{1, 3, 5, 4, 7},
		3,
	},

	{
		[]int{2, 2, 2, 2, 2},
		1,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findLengthOfLCIS(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findLengthOfLCIS(tc.nums)
		}
	}
}
