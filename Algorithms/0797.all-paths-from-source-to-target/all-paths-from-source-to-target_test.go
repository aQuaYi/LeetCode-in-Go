package problem0797

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	graph [][]int
	ans   [][]int
}{

	{
		[][]int{{1, 2}, {3}, {3}, {}},
		[][]int{{0, 1, 3}, {0, 2, 3}},
	},

	// 可以有多个 testcase
}

func Test_allPathsSourceTarget(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, allPathsSourceTarget(tc.graph), "输入:%v", tc)
	}
}

func Benchmark_allPathsSourceTarget(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			allPathsSourceTarget(tc.graph)
		}
	}
}
