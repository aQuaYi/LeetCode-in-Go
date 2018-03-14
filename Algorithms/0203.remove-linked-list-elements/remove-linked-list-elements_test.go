package problem0203

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	ints []int
	val  int
	ans  []int
}{

	{
		[]int{1, 2, 6, 3, 4, 5, 6},
		6,
		[]int{1, 2, 3, 4, 5},
	},

	// 可以有多个 testcase
}

func Test_removeElements(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		head := kit.Ints2List(tc.ints)
		ast.Equal(tc.ans, kit.List2Ints(removeElements(head, tc.val)), "输入:%v", tc)
	}
}

func Benchmark_removeElements(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			head := kit.Ints2List(tc.ints)
			removeElements(head, tc.val)
		}
	}
}
