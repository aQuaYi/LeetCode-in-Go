package problem0757

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	intervals [][]int
	ans       int
}{

	{
		[][]int{{1, 5}, {4, 5}, {5, 9}, {7, 9}, {9, 20}, {12, 20}, {15, 20}, {18, 20}},
		6,
	},

	{
		[][]int{{1, 5}, {4, 5}, {5, 9}, {7, 9}, {9, 10}},
		5,
	},

	{
		[][]int{{1, 3}, {4, 9}, {0, 10}, {6, 7}, {1, 2}, {0, 6}, {7, 9}, {0, 1}, {2, 5}, {6, 8}},
		7,
	},

	{
		[][]int{{4, 14}, {6, 17}, {7, 14}, {14, 21}, {4, 7}},
		4,
	},

	{
		[][]int{{2, 10}, {3, 7}, {3, 15}, {4, 11}, {6, 12}, {6, 16}, {7, 8}, {7, 11}, {7, 15}, {11, 12}},
		5,
	},

	{
		[][]int{{100, 101}, {1, 3}, {8, 9}},
		6,
	},

	{
		[][]int{{1, 3}, {8, 9}},
		4,
	},

	{
		[][]int{{1, 3}, {1, 4}, {2, 5}, {3, 5}},
		3,
	},

	{
		[][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}},
		5,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, intersectionSizeTwo(tc.intervals), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			intersectionSizeTwo(tc.intervals)
		}
	}
}
