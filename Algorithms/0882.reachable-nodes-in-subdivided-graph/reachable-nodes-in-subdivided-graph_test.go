package problem0882

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	edges [][]int
	M     int
	N     int
	ans   int
}{

	{
		[][]int{{1, 2, 5}, {0, 3, 3}, {1, 3, 2}, {2, 3, 4}, {0, 4, 1}},
		7,
		5,
		13,
	},

	{
		[][]int{{1, 2, 4}, {1, 4, 5}, {1, 3, 1}, {2, 3, 4}, {3, 4, 5}},
		17,
		5,
		1,
	},

	{
		[][]int{{0, 1, 10}, {0, 2, 1}, {1, 2, 2}},
		6,
		3,
		13,
	},

	{
		[][]int{{0, 1, 4}, {1, 2, 6}, {0, 2, 8}, {1, 3, 1}},
		10,
		4,
		23,
	},

	// 可以有多个 testcase
}

func Test_reachableNodes(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, reachableNodes(tc.edges, tc.M, tc.N), "输入:%v", tc)
	}
}

func Benchmark_reachableNodes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			reachableNodes(tc.edges, tc.M, tc.N)
		}
	}
}
