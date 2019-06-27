package problem1030

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	R   int
	C   int
	r0  int
	c0  int
	exp [][]int
}{

	{
		10,
		10,
		5,
		9,
		[][]int{{5, 9}, {4, 9}, {5, 8}, {6, 9}, {3, 9}, {4, 8}, {5, 7}, {6, 8}, {7, 9}, {2, 9}, {3, 8}, {4, 7}, {5, 6}, {6, 7}, {7, 8}, {8, 9}, {1, 9}, {2, 8}, {3, 7}, {4, 6}, {5, 5}, {6, 6}, {7, 7}, {8, 8}, {9, 9}, {0, 9}, {1, 8}, {2, 7}, {3, 6}, {4, 5}, {5, 4}, {6, 5}, {7, 6}, {8, 7}, {9, 8}, {0, 8}, {1, 7}, {2, 6}, {3, 5}, {4, 4}, {5, 3}, {6, 4}, {7, 5}, {8, 6}, {9, 7}, {0, 7}, {1, 6}, {2, 5}, {3, 4}, {4, 3}, {5, 2}, {6, 3}, {7, 4}, {8, 5}, {9, 6}, {0, 6}, {1, 5}, {2, 4}, {3, 3}, {4, 2}, {5, 1}, {6, 2}, {7, 3}, {8, 4}, {9, 5}, {0, 5}, {1, 4}, {2, 3}, {3, 2}, {4, 1}, {5, 0}, {6, 1}, {7, 2}, {8, 3}, {9, 4}, {0, 4}, {1, 3}, {2, 2}, {3, 1}, {4, 0}, {6, 0}, {7, 1}, {8, 2}, {9, 3}, {0, 3}, {1, 2}, {2, 1}, {3, 0}, {7, 0}, {8, 1}, {9, 2}, {0, 2}, {1, 1}, {2, 0}, {8, 0}, {9, 1}, {0, 1}, {1, 0}, {9, 0}, {0, 0}},
	},

	{
		2,
		2,
		0,
		1,
		[][]int{{0, 1}, {0, 0}, {1, 1}, {1, 0}},
	},

	{
		2,
		3,
		1,
		2,
		[][]int{{1, 2}, {0, 2}, {1, 1}, {0, 1}, {1, 0}, {0, 0}},
	},

	{
		1,
		2,
		0,
		0,
		[][]int{{0, 0}, {0, 1}},
	},
	{
		2,
		2,
		0,
		1,
		[][]int{{0, 1}, {0, 0}, {1, 1}, {1, 0}},
	},
	{
		2,
		3,
		1,
		2,
		[][]int{{1, 2}, {0, 2}, {1, 1}, {0, 1}, {1, 0}, {0, 0}},
	},

	// 可以有多个 testcase
}

func Test_allCellsDistOrder(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ans := allCellsDistOrder(tc.R, tc.C, tc.r0, tc.c0)
		ast.Equal(tc.exp, ans, "输入:%v", tc)
	}
}

func Benchmark_allCellsDistOrder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			allCellsDistOrder(tc.R, tc.C, tc.r0, tc.c0)
		}
	}
}

var limit = 10000

func Benchmark_abs_normal(b *testing.B) {
	for i := 1; i < b.N; i++ {
		for j := -limit; j <= limit; j++ {
			abs(j)
		}
	}
}

func absBit(n int) int {
	x := n >> 63
	return (n ^ x) - x
}

func Benchmark_absBit(b *testing.B) {
	for i := 1; i < b.N; i++ {
		for j := -limit; j <= limit; j++ {
			absBit(j)
		}
	}
}
