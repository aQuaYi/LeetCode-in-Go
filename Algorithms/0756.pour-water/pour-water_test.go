package problem0756

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	heights []int
	V, K    int
	ans     []int
}{

	{
		[]int{1, 2, 3, 4, 3, 2, 1, 2, 3, 4, 3, 2, 1},
		10,
		5,
		[]int{3, 3, 4, 4, 4, 4, 3, 3, 3, 4, 3, 2, 1},
	},

	{
		[]int{2, 1, 1, 2, 1, 2, 2},
		4,
		3,
		[]int{2, 2, 2, 3, 2, 2, 2},
	},

	{
		[]int{1, 2, 3, 4},
		2,
		2,
		[]int{2, 3, 3, 4},
	},

	{
		[]int{3, 1, 3},
		5,
		1,
		[]int{4, 4, 4},
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, pourWater(tc.heights, tc.V, tc.K), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			pourWater(tc.heights, tc.V, tc.K)
		}
	}
}
