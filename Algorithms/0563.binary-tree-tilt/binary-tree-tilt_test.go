package problem0563

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
		[]int{1, 2, 4, 3, 5},
		[]int{4, 2, 1, 5, 3},
		11,
	},

	{
		[]int{1, 2, 3},
		[]int{2, 1, 3},
		1,
	},

	// 可以有多个 testcase
}

func Test_findTilt(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := kit.PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.ans, findTilt(root), "输入:%v", tc)
	}
}

func Benchmark_findTilt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.PreIn2Tree(tc.pre, tc.in)
			findTilt(root)
		}
	}
}
