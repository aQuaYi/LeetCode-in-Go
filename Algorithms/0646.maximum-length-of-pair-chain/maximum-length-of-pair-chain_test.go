package problem0646

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	pairs [][]int
	ans   int
}{

	{
		[][]int{{-10, -8}, {8, 9}, {-5, 0}, {6, 10}, {-6, -4}, {1, 7}, {9, 10}, {-4, 7}},
		4,
	},

	{
		[][]int{{1, 2}, {2, 3}, {3, 4}},
		2,
	},

	// 可以有多个 testcase
}

func Test_fcName(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findLongestChain(tc.pairs), "输入:%v", tc)
	}
}

func Benchmark_fcName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findLongestChain(tc.pairs)
		}
	}
}
