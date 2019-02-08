package problem0956

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	rods []int
	ans  int
}{

	{
		[]int{1, 2, 3, 6},
		6,
	},

	{
		[]int{1, 2, 3, 4, 5, 6},
		10,
	},

	{
		[]int{1, 2},
		0,
	},

	// 可以有多个 testcase
}

func Test_tallestBillboard(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, tallestBillboard(tc.rods), "输入:%v", tc)
	}
}

func Benchmark_tallestBillboard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			tallestBillboard(tc.rods)
		}
	}
}
