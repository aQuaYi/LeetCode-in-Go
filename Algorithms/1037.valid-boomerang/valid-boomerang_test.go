package problem1037

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	points [][]int
	ans    bool
}{

	{
		[][]int{{1, 1}, {2, 3}, {3, 2}},
		true,
	},

	{
		[][]int{{1, 1}, {2, 2}, {3, 3}},
		false,
	},

	// 可以有多个 testcase
}

func Test_isBoomerang(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, isBoomerang(tc.points), "输入:%v", tc)
	}
}

func Benchmark_isBoomerang(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isBoomerang(tc.points)
		}
	}
}
