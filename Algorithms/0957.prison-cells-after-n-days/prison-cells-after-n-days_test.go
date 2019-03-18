package problem0957

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	cells []int
	N     int
	ans   []int
}{

	{
		[]int{1, 0, 0, 1, 0, 0, 0, 1},
		826,
		[]int{0, 1, 1, 0, 1, 1, 1, 0},
	},

	{
		[]int{0, 1, 0, 1, 1, 0, 0, 1},
		1,
		[]int{0, 1, 1, 0, 0, 0, 0, 0},
	},

	{
		[]int{0, 1, 0, 1, 1, 0, 0, 1},
		7,
		[]int{0, 0, 1, 1, 0, 0, 0, 0},
	},

	{
		[]int{1, 0, 0, 1, 0, 0, 1, 0},
		1000000000,
		[]int{0, 0, 1, 1, 1, 1, 1, 0},
	},

	// 可以有多个 testcase
}

func Test_prisonAfterNDays(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, prisonAfterNDays(tc.cells, tc.N), "输入:%v", tc)
	}
}

func Benchmark_prisonAfterNDays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			prisonAfterNDays(tc.cells, tc.N)
		}
	}
}
