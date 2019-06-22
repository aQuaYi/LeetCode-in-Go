package problem1036

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	blocked [][]int
	source  []int
	target  []int
	ans     bool
}{

	{
		[][]int{{0, 1}, {1, 0}},
		[]int{0, 0},
		[]int{0, 2},
		false,
	},

	{
		[][]int{},
		[]int{0, 0},
		[]int{999999, 999999},
		true,
	},

	// 可以有多个 testcase
}

func Test_isEscapePossible(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, isEscapePossible(tc.blocked, tc.source, tc.target), "输入:%v", tc)
	}
}

func Benchmark_isEscapePossible(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isEscapePossible(tc.blocked, tc.source, tc.target)
		}
	}
}
