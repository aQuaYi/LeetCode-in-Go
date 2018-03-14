package problem0671

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	in, post []int
	ans      int
}{

	{
		[]int{2, 2, 2},
		[]int{2, 2, 2},
		-1,
	},

	{
		[]int{4, 2, 5, 1, 6, 3, 7},
		[]int{4, 5, 2, 6, 7, 3, 1},
		2,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := kit.InPost2Tree(tc.in, tc.post)
		ast.Equal(tc.ans, findSecondMinimumValue(root), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.PreIn2Tree(tc.in, tc.post)
			findSecondMinimumValue(root)
		}
	}
}
