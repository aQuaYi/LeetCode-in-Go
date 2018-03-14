package problem0486

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
		[]int{0},
		true,
	},

	{
		[]int{2},
		true,
	},

	{
		[]int{1, 1},
		true,
	},

	{
		[]int{1, 5, 2},
		false,
	},

	{
		[]int{1, 5, 233, 7},
		true,
	},

	// 可以有多个 testcase
}

func Test_PredictTheWinner(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, PredictTheWinner(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_PredictTheWinner(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			PredictTheWinner(tc.nums)
		}
	}
}
