package problem0260

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  []int
}{

	{
		[]int{1, 2, 3, 1, 2, 5},
		[]int{3, 5},
	},

	{
		[]int{1, 2, -1, 1, 2, 0},
		[]int{-1, 0},
	},
	// 可以有多个 testcase
}

func Test_singleNumber(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ans := singleNumber(tc.nums)
		sort.Ints(ans)
		ast.Equal(tc.ans, ans, "输入:%v", tc)
	}
}

func Benchmark_singleNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			singleNumber(tc.nums)
		}
	}
}
