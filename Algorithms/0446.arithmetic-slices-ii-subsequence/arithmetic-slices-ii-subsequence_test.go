package Problem0446

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
		[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		134217349,
	},

	{
		[]int{2, 4, 4, 6},
		2,
	},

	{
		[]int{8, 10},
		0,
	},

	{
		[]int{2, 4, 6, 8, 10},
		7,
	},

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
