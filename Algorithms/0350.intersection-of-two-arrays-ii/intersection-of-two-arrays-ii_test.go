package problem0350

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums1 []int
	nums2 []int
	ans   []int
}{

	{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
	{[]int{1, 3, 3, 3, 2, 2, 1}, []int{3, 2, 2}, []int{2, 2, 3}},

	// 可以有多个 testcase
}

func Test_intersect(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ans := intersect(tc.nums1, tc.nums2)
		sort.Ints(ans)

		ast.Equal(tc.ans, ans, "输入:%v", tc)
	}
}

func Benchmark_intersect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			intersect(tc.nums1, tc.nums2)
		}
	}
}
