package problem0876

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	head []int
	ans  []int
}{

	{
		[]int{1, 2, 3, 4, 5},
		[]int{3, 4, 5},
	},

	{
		[]int{1, 2, 3, 4, 5, 6},
		[]int{4, 5, 6},
	},

	// 可以有多个 testcase
}

func Test_middleNode(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		head := kit.Ints2List(tc.head)
		actual := kit.List2Ints(middleNode(head))
		ast.Equal(tc.ans, actual, "输入:%v", tc)
	}
}

func Benchmark_middleNode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			head := kit.Ints2List(tc.head)
			middleNode(head)
		}
	}
}
