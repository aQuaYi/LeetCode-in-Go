package problem0026

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
		[]int{1, 1, 2},
		2,
	},

	{
		[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		5,
	},

	// 可以有多个 testcase
}

func Test_removeDuplicates(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, removeDuplicates(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_removeDuplicates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			removeDuplicates(tc.nums)
		}
	}
}
