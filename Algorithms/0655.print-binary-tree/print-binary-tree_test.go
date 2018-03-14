package problem0655

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	pre, in []int
	ans     [][]string
}{

	{
		[]int{1, 2},
		[]int{2, 1},
		[][]string{{"", "1", ""},
			{"2", "", ""}},
	},

	{
		[]int{1, 2, 4, 3},
		[]int{2, 4, 1, 3},
		[][]string{{"", "", "", "1", "", "", ""},
			{"", "2", "", "", "", "3", ""},
			{"", "", "4", "", "", "", ""}},
	},

	{
		[]int{1, 2, 3, 4, 5},
		[]int{4, 3, 2, 1, 5},
		[][]string{{"", "", "", "", "", "", "", "1", "", "", "", "", "", "", ""},
			{"", "", "", "2", "", "", "", "", "", "", "", "5", "", "", ""},
			{"", "3", "", "", "", "", "", "", "", "", "", "", "", "", ""},
			{"4", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}},
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := kit.PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.ans, printTree(root), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.PreIn2Tree(tc.pre, tc.in)
			printTree(root)
		}
	}
}
