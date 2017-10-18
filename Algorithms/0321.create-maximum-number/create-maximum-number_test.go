package Problem0321

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums1 []int
	nums2 []int
	k     int
	ans   []int
}{
	{
		[]int{3, 4, 6, 5},
		[]int{9, 1, 2, 5, 8, 3},
		5,
		[]int{9, 8, 6, 5, 3},
	},

	{

		[]int{6, 7},
		[]int{6, 0, 4},
		5,
		[]int{6, 7, 6, 0, 4},
	},

	{

		[]int{3, 9},
		[]int{8, 9},
		3,
		[]int{9, 8, 9},
	},

	// 可以有多个 testcase
}

func Test_maxNumber(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, maxNumber(tc.nums1, tc.nums2, tc.k), "输入:%v", tc)
	}
}

func Benchmark_maxNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxNumber(tc.nums1, tc.nums2, tc.k)
		}
	}
}
