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
		[][]int{{3, 4, 8}, {0, 1, 3}, {1, 4, 0}, {1, 2, 3}, {0, 3, 2}, {0, 4, 10}, {1, 3, 3}, {2, 4, 3}, {2, 3, 3}, {0, 2, 10}},
		7,
		5,
		43,
	},

	{
		[][]int{{4, 5, 21}, {2, 9, 19}, {5, 9, 12}, {3, 8, 17}, {4, 9, 2}, {1, 2, 19}, {2, 8, 6}, {8, 9, 20}, {3, 5, 16}, {5, 6, 13}, {0, 8, 13}, {5, 8, 1}, {0, 9, 1}, {6, 8, 0}, {1, 3, 20}, {7, 9, 3}, {6, 9, 3}, {1, 9, 10}, {5, 7, 25}, {3, 6, 23}, {4, 6, 8}, {4, 7, 21}, {0, 5, 2}, {0, 6, 10}, {4, 8, 15}},
		20,
		10,
		282,
	},

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
