package Problem0327

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums  []int
	lower int
	upper int
	ans   int
}{

	{[]int{-2, 5, -1}, -2, 2, 3},

	// 可以有多个 testcase
}

func Test_countRangeSum(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, countRangeSum(tc.nums, tc.lower, tc.upper), "输入:%v", tc)
	}
}

func Benchmark_countRangeSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			countRangeSum(tc.nums, tc.lower, tc.upper)
		}
	}
}
