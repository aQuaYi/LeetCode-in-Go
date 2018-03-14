package problem0147

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
	head := kit.Ints2List([]int{9, 6, 3, 1, 4, 8, 2, 5, 7, 19, 16, 13, 11, 14, 18, 12, 15, 17, 29, 26, 23, 21, 24, 28, 22, 25, 27, 39, 36, 33, 31, 34, 38, 32, 35, 37, 49, 46, 43, 41, 44, 48, 42, 45, 47})
	for i := 0; i < b.N; i++ {
		insertionSortList(head)
	}
}
