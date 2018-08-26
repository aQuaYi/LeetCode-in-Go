package problem0889

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	pre  []int
	post []int
	ans  []int
}{

	{
		[]int{1, 2, 4, 5, 3, 6},
		[]int{4, 5, 2, 6, 3, 1},
		[]int{1, 2, 3, 4, 5, 6},
	},

	{
		[]int{1, 2, 4, 5, 3, 6, 7},
		[]int{4, 5, 2, 6, 7, 3, 1},
		[]int{1, 2, 3, 4, 5, 6, 7},
	},

	// 可以有多个 testcase
}

func Test_constructFromPrePost(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ans := kit.Tree2ints(constructFromPrePost(tc.pre, tc.post))
		ast.Equal(tc.ans, ans, "输入:%v", tc)
	}
}

func Benchmark_constructFromPrePost(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			constructFromPrePost(tc.pre, tc.post)
		}
	}
}
