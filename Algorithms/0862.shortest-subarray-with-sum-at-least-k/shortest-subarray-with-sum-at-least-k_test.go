package problem0862

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	K   int
	ans int
}{

	{
		[]int{48, 99, 37, 4, -31},
		140,
		2,
	},

	{
		[]int{1},
		1,
		1,
	},

	{
		[]int{1, 2},
		4,
		-1,
	},

	{
		[]int{2, -1, 2},
		3,
		3,
	},

	// 可以有多个 testcase
}

func Test_shortestSubarray(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, shortestSubarray(tc.A, tc.K), "输入:%v", tc)
	}
}

func Benchmark_shortestSubarray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			shortestSubarray(tc.A, tc.K)
		}
	}
}
