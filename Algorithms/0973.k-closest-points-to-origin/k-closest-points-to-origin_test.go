package problem0973

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	points [][]int
	K      int
	ans    [][]int
}{

	{
		[][]int{{1, 3}, {-2, 2}},
		1,
		[][]int{{-2, 2}},
	},

	{
		[][]int{{3, 3}, {5, -1}, {-2, 4}},
		2,
		[][]int{{3, 3}, {-2, 4}},
	},

	// 可以有多个 testcase
}

func Test_kClosest(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, kClosest(tc.points, tc.K), "输入:%v", tc)
	}
}

func Benchmark_kClosest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			kClosest(tc.points, tc.K)
		}
	}
}
