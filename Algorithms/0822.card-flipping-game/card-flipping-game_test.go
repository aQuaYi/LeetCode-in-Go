package problem0822

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	fronts []int
	backs  []int
	ans    int
}{

	{
		[]int{1, 1},
		[]int{2, 2},
		1,
	},

	{
		[]int{1, 3, 4, 4, 3},
		[]int{1, 3, 4, 1, 3},
		0,
	},

	{
		[]int{1, 2, 4, 4, 7},
		[]int{1, 3, 4, 1, 3},
		2,
	},

	// 可以有多个 testcase
}

func Test_flipgame(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, flipgame(tc.fronts, tc.backs), "输入:%v", tc)
	}
}

func Benchmark_flipgame(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			flipgame(tc.fronts, tc.backs)
		}
	}
}
