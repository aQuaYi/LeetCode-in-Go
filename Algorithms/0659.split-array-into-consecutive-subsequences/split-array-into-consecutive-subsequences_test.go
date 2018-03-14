package problem0659

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
		[]int{1, 2, 3, 3, 4, 5},
		true,
	},

	{
		[]int{1, 2, 2, 3, 3, 3, 4, 4, 5},
		true,
	},

	{
		[]int{1, 2, 3, 3, 4, 4, 5, 5},
		true,
	},

	{
		[]int{1, 2, 3, 4, 4, 5},
		false,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, isPossible(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isPossible(tc.nums)
		}
	}
}
