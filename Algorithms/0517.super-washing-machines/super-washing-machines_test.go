package problem0517

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	machines []int
	ans      int
}{

	{
		[]int{6, 3, 6},
		1,
	},

	{
		[]int{4, 0, 0, 4},
		2,
	},

	{
		[]int{1, 0, 5, 7, 8, 9},
		9,
	},

	{
		[]int{1, 0, 5},
		3,
	},

	{
		[]int{0, 3, 0},
		2,
	},

	{
		[]int{0, 2, 0},
		-1,
	},

	// 可以有多个 testcase
}

func Test_findMinMoves(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findMinMoves(tc.machines), "输入:%v", tc)
	}
}

func Benchmark_findMinMoves(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findMinMoves(tc.machines)
		}
	}
}
