package problem0347

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	k    int
	ans  []int
}{

	{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},

	// 可以有多个 testcase
}

func Test_topKFrequent(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ans := topKFrequent(tc.nums, tc.k)
		sort.Ints(ans)
		ast.Equal(tc.ans, ans, "输入:%v", tc)
	}
}

func Benchmark_topKFrequent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			topKFrequent(tc.nums, tc.k)
		}
	}
}
