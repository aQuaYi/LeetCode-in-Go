package problem0665

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
		[]int{2, 3, 3, 2, 4},
		true,
	},

	{
		[]int{3, 4, 2, 3},
		false,
	},

	{
		[]int{1, 2, 3},
		true,
	},

	{
		[]int{4, 2, 3},
		true,
	},

	{
		[]int{4, 2, 1},
		false,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, checkPossibility(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			checkPossibility(tc.nums)
		}
	}
}
