package problem0847

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	graph [][]int
	ans   int
}{

	{
		[][]int{{1}, {2}, {0}},
		2,
	},

	{
		[][]int{{1, 2, 3}, {0}, {0}, {0}},
		4,
	},

	{
		[][]int{{1}, {0, 2, 4}, {1, 3, 4}, {2}, {1, 2}},
		4,
	},

	// 可以有多个 testcase
}

func Test_shortestPathLength(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, shortestPathLength(tc.graph), "输入:%v", tc)
	}
}

func Benchmark_shortestPathLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			shortestPathLength(tc.graph)
		}
	}
}
