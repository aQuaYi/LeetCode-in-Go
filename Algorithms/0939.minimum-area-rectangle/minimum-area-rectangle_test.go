package problem0939

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	points [][]int
	ans    int
}{

	{
		[][]int{{3, 2}, {1, 3}, {3, 0}, {3, 4}, {2, 1}, {0, 4}, {0, 3}, {4, 1}, {2, 4}},
		0,
	},

	{
		[][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}, {2, 2}},
		4,
	},

	{
		[][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}, {4, 1}, {4, 3}},
		2,
	},

	{
		[][]int{{3, 2}, {3, 1}, {4, 4}, {1, 1}, {4, 3}, {0, 3}, {0, 2}, {4, 0}},
		0,
	},

	// 可以有多个 testcase
}

func Test_minAreaRect(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Println("====")
		ast.Equal(tc.ans, minAreaRect(tc.points), "输入:%v", tc)
	}
}

func Benchmark_minAreaRect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minAreaRect(tc.points)
		}
	}
}
