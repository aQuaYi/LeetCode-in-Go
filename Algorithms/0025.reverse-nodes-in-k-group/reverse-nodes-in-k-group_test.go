package problem0025

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	head []int
	k    int
	ans  []int
}{

	{
		[]int{1, 2, 3, 4, 5},
		3,
		[]int{3, 2, 1, 4, 5},
	},

	{
		[]int{1, 2, 3, 4, 5},
		1,
		[]int{1, 2, 3, 4, 5},
	},

	// 可以有多个 testcase
}

func Test_reverseKGroup(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		head := kit.Ints2List(tc.head)
		ans := kit.Ints2List(tc.ans)
		ast.Equal(ans, reverseKGroup(head, tc.k), "输入:%v", tc)
	}
}

func Benchmark_reverseKGroup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			head := kit.Ints2List(tc.head)
			reverseKGroup(head, tc.k)
		}
	}
}
