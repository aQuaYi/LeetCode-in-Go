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
		ast.Equal(len(tc.exp), len(ans), "输入:%v", tc)
		n := len(ans)
		for i := 0; i < n; i++ {
			e, a := tc.exp[i], ans[i]
			ast.Equal(dist(e[0], e[1], tc.r0, tc.c0), dist(a[0], a[1], tc.r0, tc.c0), "%v,%v, %v", e, a, tc)
		}
	}
}

func Benchmark_allCellsDistOrder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			allCellsDistOrder(tc.R, tc.C, tc.r0, tc.c0)
		}
	}
}

func dist(r1, c1, r2, c2 int) int {
	return abs(r1-r2) + abs(c1-c2)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
