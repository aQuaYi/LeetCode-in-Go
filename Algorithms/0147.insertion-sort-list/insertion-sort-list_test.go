package Problem0147

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
		[]int{2, 4, 3, 1, 5},
		[]int{1, 2, 3, 4, 5},
	},

	// 可以有多个 testcase
}

func Test_insertionSortList(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		head := kit.Ints2List(tc.head)
		ast.Equal(tc.ans, kit.List2Ints(insertionSortList(head)), "输入:%v", tc)
	}
}

func Benchmark_insertionSortList(b *testing.B) {
	head := &ListNode{}
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			b.StopTimer()
			head = kit.Ints2List(tc.head)
			b.StartTimer()
			insertionSortList(head)
		}
	}
}
