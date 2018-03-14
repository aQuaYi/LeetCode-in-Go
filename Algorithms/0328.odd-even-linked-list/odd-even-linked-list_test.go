package problem0328

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
		[]int{1, 3, 5, 2, 4},
	},

	{
		[]int{1, 2, 3, 4, 5, 6},
		[]int{1, 3, 5, 2, 4, 6},
	},

	{
		[]int{},
		[]int{},
	},
	// 可以有多个 testcase
}

func Test_oddEvenList(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		head := kit.Ints2List(tc.head)
		ast.Equal(tc.ans, kit.List2Ints(oddEvenList(head)), "输入:%v", tc)
	}
}

func Benchmark_oddEvenList(b *testing.B) {
	head := kit.Ints2List([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	for i := 0; i < b.N; i++ {
		oddEvenList(head)
	}
}
