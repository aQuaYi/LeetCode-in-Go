package problem0740

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
		[]int{8, 3, 4, 7, 6, 6, 9, 2, 5, 8, 2, 4, 9, 5, 9, 1, 5, 7, 1, 4},
		61,
	},

	{
		[]int{},
		0,
	},

	{
		[]int{3},
		3,
	},

	{
		[]int{3, 4, 2},
		6,
	},

	{
		[]int{2, 2, 3, 3, 3, 4},
		9,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, deleteAndEarn(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			deleteAndEarn(tc.nums)
		}
	}
}
