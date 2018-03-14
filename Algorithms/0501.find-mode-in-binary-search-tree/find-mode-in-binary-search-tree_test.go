package problem0501

import (
	"fmt"
	"sort"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	pre, in []int
	ans     []int
}{

	{
		[]int{1, 2, 1, 2},
		[]int{1, 1, 2, 2},
		[]int{1, 2},
	},

	// 可以有多个 testcase
}

func Test_findMode(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := kit.PreIn2Tree(tc.pre, tc.in)
		ans := findMode(root)
		sort.Ints(ans)
		ast.Equal(tc.ans, ans, "输入:%v", tc)
	}
}

func Benchmark_findMode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.PreIn2Tree(tc.pre, tc.in)
			findMode(root)
		}
	}
}
