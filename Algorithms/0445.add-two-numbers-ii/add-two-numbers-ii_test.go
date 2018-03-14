package problem0445

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	l1  []int
	l2  []int
	ans []int
}{

	{
		[]int{7, 2, 4, 3},
		[]int{4, 5, 6, 4},
		[]int{1, 1, 8, 0, 7},
	},

	{
		[]int{7, 2, 4, 3},
		[]int{1, 5, 6, 4},
		[]int{8, 8, 0, 7},
	},

	{
		[]int{7, 2, 4, 3},
		[]int{5, 6, 4},
		[]int{7, 8, 0, 7},
	},

	// 可以有多个 testcase
}

func Test_addTwoNumbers(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		l1 := kit.Ints2List(tc.l1)
		l2 := kit.Ints2List(tc.l2)
		ast.Equal(tc.ans, kit.List2Ints(addTwoNumbers(l1, l2)), "输入:%v", tc)
	}
}

func Benchmark_addTwoNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			l1 := kit.Ints2List(tc.l1)
			l2 := kit.Ints2List(tc.l2)
			addTwoNumbers(l1, l2)
		}
	}
}
