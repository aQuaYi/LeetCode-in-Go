package problem0508

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
		[]int{5, 2, -5},
		[]int{2, 5, -5},
		[]int{2},
	},

	{
		[]int{5, 2, -3},
		[]int{2, 5, -3},
		[]int{2, -3, 4},
	},

	// 可以有多个 testcase
}

func Test_findFrequentTreeSum(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := kit.PreIn2Tree(tc.pre, tc.in)
		ans := findFrequentTreeSum(root)
		sort.Ints(ans)
		sort.Ints(tc.ans)
		ast.Equal(tc.ans, ans, "输入:%v", tc)
	}
}

func Benchmark_findFrequentTreeSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.PreIn2Tree(tc.pre, tc.in)
			findFrequentTreeSum(root)
		}
	}
}
