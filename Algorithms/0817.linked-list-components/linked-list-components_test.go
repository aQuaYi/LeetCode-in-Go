package problem0817

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	head []int
	G    []int
	ans  int
}{

	{
		[]int{0, 2, 4, 3, 1},
		[]int{3, 2, 4},
		1,
	},

	{
		[]int{1, 2, 0, 4, 3},
		[]int{3, 4, 0, 2, 1},
		1,
	},

	{
		[]int{0, 1, 2},
		[]int{1, 0, 2},
		1,
	},

	{
		[]int{0, 1, 2},
		[]int{1, 0},
		1,
	},

	{
		[]int{0, 1, 2, 3},
		[]int{0, 1, 3},
		2,
	},

	{
		[]int{0, 1, 2, 3, 4},
		[]int{0, 3, 1, 4},
		2,
	},

	// 可以有多个 testcase
}

func Test_numComponents(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		head := kit.Ints2List(tc.head)
		ast.Equal(tc.ans, numComponents(head, tc.G), "输入:%v", tc)
	}
}

func Benchmark_numComponents(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			head := kit.Ints2List(tc.head)
			numComponents(head, tc.G)
		}
	}
}
