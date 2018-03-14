package problem0679

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
		[]int{1, 2, 1, 2},
		false,
	},

	{
		[]int{4, 1, 8, 7},
		true,
	},

	{
		[]int{1, 5, 5, 5},
		true,
	},

	{
		[]int{3, 8, 8, 8},
		true,
	},

	// 可以有多个 testcase
}

func Test_judgePoint24(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, judgePoint24(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_judgePoint24(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			judgePoint24(tc.nums)
		}
	}
}
