package problem0769

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	arr []int
	ans int
}{

	{
		[]int{2, 1, 0, 5, 4, 3, 6, 9, 8, 7},
		4,
	},

	{
		[]int{4, 3, 2, 1, 0},
		1,
	},

	{
		[]int{1, 0, 2, 3, 4},
		4,
	},

	// 可以有多个 testcase
}

func Test_maxChunksToSorted(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, maxChunksToSorted(tc.arr), "输入:%v", tc)
	}
}

func Benchmark_maxChunksToSorted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxChunksToSorted(tc.arr)
		}
	}
}
