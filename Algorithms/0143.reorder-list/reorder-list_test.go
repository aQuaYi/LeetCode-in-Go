package Problem0143

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Golang/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	head []int
	ans  []int
}{

	{
		[]int{1, 2, 3, 4},
		[]int{1, 4, 2, 3},
	},

	// 可以有多个 testcase
}

func Test_reorderList(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		head := kit.Slice2List(tc.head)
		reorderList(head)
		ans := kit.List2Slice(head)

		ast.Equal(tc.ans, ans, "输入:%v", tc)
	}
}

func Benchmark_reorderList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			reorderList(kit.Slice2List(tc.head))
		}
	}
}
