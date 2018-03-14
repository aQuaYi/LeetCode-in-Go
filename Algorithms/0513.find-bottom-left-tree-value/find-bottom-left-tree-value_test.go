package problem0513

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	pre, in []int
	ans     int
}{

	{
		[]int{0, -1},
		[]int{0, -1},
		-1,
	},

	{
		[]int{1},
		[]int{1},
		1,
	},

	{
		[]int{2, 1, 3},
		[]int{1, 2, 3},
		1,
	},

	{
		[]int{1, 2, 4, 3, 5, 6},
		[]int{4, 2, 1, 5, 3, 6},
		4,
	},

	{
		[]int{1, 2, 4, 3, 5, 7, 6},
		[]int{4, 2, 1, 7, 5, 3, 6},
		7,
	},

	// 可以有多个 testcase
}

func Test_findBottomLeftValue(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := kit.PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.ans, findBottomLeftValue(root), "输入:%v", tc)
	}
}

func Benchmark_findBottomLeftValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.PreIn2Tree(tc.pre, tc.in)
			findBottomLeftValue(root)
		}
	}
}
