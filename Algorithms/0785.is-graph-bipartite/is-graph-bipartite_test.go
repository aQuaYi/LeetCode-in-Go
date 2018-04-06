package problem0785

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	graph [][]int
	ans   bool
}{
	{
		[][]int{{}, {2, 4, 6}, {1, 4, 8, 9}, {7, 8}, {1, 2, 8, 9}, {6, 9}, {1, 5, 7, 8, 9}, {3, 6, 9}, {2, 3, 4, 6, 9}, {2, 4, 5, 6, 7, 8}},
		false,
	},
	{
		[][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}},
		true,
	},

	{
		[][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}},
		false,
	},

	// 可以有多个 testcase
}

func Test_isBipartite(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, isBipartite(tc.graph), "输入:%v", tc)
	}
}

func Benchmark_isBipartite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isBipartite(tc.graph)
		}
	}
}
