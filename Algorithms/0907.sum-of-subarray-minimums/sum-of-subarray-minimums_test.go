package problem0907

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

	{
		[]int{48, 87, 27},
		264,
	},

	{
		[]int{3, 1, 2, 4},
		17,
	},

	// 可以有多个 testcase
}

func Test_sumSubarrayMins(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, sumSubarrayMins(tc.A), "输入:%v", tc)
	}
}

func Benchmark_sumSubarrayMins(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			sumSubarrayMins(tc.A)
		}
	}
}
