package problem0349

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums1 []int
	nums2 []int
	ans   []int
}{

	{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2}},

	// 可以有多个 testcase
}

func Test_intersection(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, intersection(tc.nums1, tc.nums2), "输入:%v", tc)
	}
}

func Benchmark_intersection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			intersection(tc.nums1, tc.nums2)
		}
	}
}
