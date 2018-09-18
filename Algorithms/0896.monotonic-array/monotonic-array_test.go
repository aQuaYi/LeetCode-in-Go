package problem0896

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	ans bool
}{

	{
		[]int{1, 2, 2, 3},
		true,
	},

	{
		[]int{6, 5, 4, 4},
		true,
	},

	{
		[]int{1, 3, 2},
		false,
	},

	{
		[]int{1, 2, 4, 5},
		true,
	},

	{
		[]int{1, 1, 1},
		true,
	},

	// 可以有多个 testcase
}

func Test_isMonotonic(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, isMonotonic(tc.A), "输入:%v", tc)
	}
}

func Benchmark_isMonotonic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isMonotonic(tc.A)
		}
	}
}
