package problem0834

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	N     int
	edges [][]int
	ans   []int
}{

	{
		6,
		[][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}},
		[]int{8, 12, 6, 10, 10, 10},
	},

	// 可以有多个 testcase
}

func Test_sumOfDistancesInTree(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, sumOfDistancesInTree(tc.N, tc.edges), "输入:%v", tc)
	}
}

func Benchmark_sumOfDistancesInTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			sumOfDistancesInTree(tc.N, tc.edges)
		}
	}
}
