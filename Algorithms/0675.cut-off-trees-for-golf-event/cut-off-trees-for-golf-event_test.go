package Problem0675

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	forest [][]int
	ans    int
}{

	{
		[][]int{{1, 2, 3}, {0, 0, 4}, {7, 6, 5}},
		6,
	},

	{
		[][]int{{1, 2, 3}, {0, 0, 0}, {7, 6, 5}},
		-1,
	},

	{
		[][]int{{2, 3, 4}, {0, 0, 5}, {8, 7, 6}},
		6,
	},

	// 可以有多个 testcase
}

func Test_cutOffTree(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, cutOffTree(tc.forest), "输入:%v", tc)
	}
}

func Benchmark_cutOffTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			cutOffTree(tc.forest)
		}
	}
}
