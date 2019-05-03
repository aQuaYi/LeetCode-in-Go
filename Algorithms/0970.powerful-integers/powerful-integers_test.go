package problem0970

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	x     int
	y     int
	bound int
	ans   []int
}{

	{
		2,
		3,
		10,
		[]int{2, 3, 4, 5, 7, 9, 10},
	},

	{
		3,
		5,
		15,
		[]int{2, 4, 6, 8, 10, 14},
	},

	// 可以有多个 testcase
}

func Test_powerfulIntegers(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, powerfulIntegers(tc.x, tc.y, tc.bound), "输入:%v", tc)
	}
}

func Benchmark_powerfulIntegers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			powerfulIntegers(tc.x, tc.y, tc.bound)
		}
	}
}
