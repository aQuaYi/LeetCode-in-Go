package problem0725

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	ints []int
	k    int
	ans  [][]int
}{

	{
		[]int{1, 2, 3},
		5,
		[][]int{{1}, {2}, {3}, {}, {}},
	},

	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		3,
		[][]int{{1, 2, 3, 4}, {5, 6, 7}, {8, 9, 10}},
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := kit.Ints2List(tc.ints)
		ansLN := splitListToParts(root, tc.k)
		ans := make([][]int, len(tc.ans))
		for i := range ans {
			ans[i] = kit.List2Ints(ansLN[i])
		}
		ast.Equal(tc.ans, ans, "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.Ints2List(tc.ints)
			splitListToParts(root, tc.k)
		}
	}
}
