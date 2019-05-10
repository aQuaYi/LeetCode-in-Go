package problem0997

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	N     int
	trust [][]int
	ans   int
}{

	{
		1,
		[][]int{},
		1,
	},

	{
		2,
		[][]int{{1, 2}},
		2,
	},

	{
		3,
		[][]int{{1, 3}, {2, 3}},
		3,
	},

	{
		3,
		[][]int{{1, 3}, {2, 3}, {3, 1}},
		-1,
	},

	{
		3,
		[][]int{{1, 2}, {2, 3}},
		-1,
	},

	{
		4,
		[][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}},
		3,
	},

	// 可以有多个 testcase
}

func Test_findJudge(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, findJudge(tc.N, tc.trust), "输入:%v", tc)
	}
}

func Benchmark_findJudge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findJudge(tc.N, tc.trust)
		}
	}
}
