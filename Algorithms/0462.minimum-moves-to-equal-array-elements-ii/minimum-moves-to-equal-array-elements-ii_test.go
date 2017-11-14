package Problem0462

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  int
}{

	{
		[]int{1, 0, 0, 8, 6},
		14,
	},

	{
		[]int{3, 2, 2, 2, 2, 3, 3, 3, 3},
		4,
	},

	{
		[]int{2, 2, 2, 2, 2, 3, 3, 3, 3},
		4,
	},

	{
		[]int{1, 2, 3},
		2,
	},

	// 可以有多个 testcase
}

func Test_minMoves2(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, minMoves2(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_minMoves2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minMoves2(tc.nums)
		}
	}
}
