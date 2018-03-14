package problem0768

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
		[]int{6, 4, 5},
		1,
	},

	{
		[]int{6, 4, 3, 2, 1, 5},
		1,
	},

	{
		[]int{5, 4, 3, 2, 1},
		1,
	},

	{
		[]int{2, 1, 3, 4, 4},
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
