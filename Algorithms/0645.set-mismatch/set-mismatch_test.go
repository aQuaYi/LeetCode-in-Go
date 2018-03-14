package problem0645

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  []int
}{

	{
		[]int{1, 1},
		[]int{1, 2},
	},

	{
		[]int{1, 3, 4, 5, 5, 6},
		[]int{5, 2},
	},

	{
		[]int{1, 1, 2, 3, 4, 6},
		[]int{1, 5},
	},

	{
		[]int{1, 2, 2, 4},
		[]int{2, 3},
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findErrorNums(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findErrorNums(tc.nums)
		}
	}
}
