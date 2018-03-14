package problem0540

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
		[]int{1, 1, 2, 3, 3, 4, 4, 8, 8},
		2,
	},

	{
		[]int{3, 3, 7, 7, 10, 11, 11},
		10,
	},

	// 可以有多个 testcase
}

func Test_singleNonDuplicate(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, singleNonDuplicate(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_singleNonDuplicate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			singleNonDuplicate(tc.nums)
		}
	}
}
