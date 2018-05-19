package problem0823

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	ans int
}{

	{
		[]int{18, 3, 6, 2},
		12,
	},

	{
		[]int{2, 4},
		3,
	},

	{
		[]int{2, 4, 5, 10},
		7,
	},

	// 可以有多个 testcase
}

func Test_numFactoredBinaryTrees(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, numFactoredBinaryTrees(tc.A), "输入:%v", tc)
	}
}

func Benchmark_numFactoredBinaryTrees(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numFactoredBinaryTrees(tc.A)
		}
	}
}
