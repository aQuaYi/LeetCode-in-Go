package problem0413

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	ans int
}{

	{[]int{}, 0},
	{[]int{1, 2, 3, 4}, 3},
	{[]int{2, 1, 2, 3, 4}, 3},
	{[]int{1, 2, 3, 4, 4}, 3},
	{[]int{1, 1, 1, 2, 2, 2}, 2},

	// 可以有多个 testcase
}

func Test_numberOfArithmeticSlices(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, numberOfArithmeticSlices(tc.A), "输入:%v", tc)
	}
}

func Benchmark_numberOfArithmeticSlices(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numberOfArithmeticSlices(tc.A)
		}
	}
}
