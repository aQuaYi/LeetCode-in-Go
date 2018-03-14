package problem0623

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	pre, in []int
	v, d    int
	ans     []int
}{

	{
		[]int{4, 2, 3, 1, 6, 5},
		[]int{3, 2, 1, 4, 5, 6},
		1,
		2,
		[]int{3, 1, 2, 1, 5, 6, 1, 4},
	},

	{
		[]int{1, 2, 3},
		[]int{2, 1, 3},
		1,
		1,
		[]int{2, 3, 1, 1},
	},

	{
		[]int{1, 2, 3, 4, 6},
		[]int{1, 3, 2, 4, 6},
		5,
		4,
		[]int{5, 5, 3, 5, 6, 5, 4, 2, 1},
	},

	{
		[]int{4, 2, 3, 1},
		[]int{3, 2, 1, 4},
		1,
		3,
		[]int{3, 1, 1, 1, 2, 4},
	},

	// 可以有多个 testcase
}

func Test_fcName(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := kit.PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.ans, kit.Tree2Postorder(addOneRow(root, tc.v, tc.d)), "输入:%v", tc)
	}
}

func Benchmark_fcName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.PreIn2Tree(tc.pre, tc.in)
			addOneRow(root, tc.v, tc.d)
		}
	}
}
