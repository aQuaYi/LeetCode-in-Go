package Problem0220

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	k    int
	t    int
	ans  bool
}{

	{
		[]int{2, 4, 3, 8},
		2,
		1,
		true,
	},

	{
		[]int{2, 4, 6, 8},
		1,
		1,
		false,
	},

	// 可以有多个 testcase
}

func Test_containsNearbyAlmostDuplicate(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, containsNearbyAlmostDuplicate(tc.nums, tc.k, tc.t), "输入:%v", tc)
	}
}

func Benchmark_containsNearbyAlmostDuplicate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			containsNearbyAlmostDuplicate(tc.nums, tc.k, tc.t)
		}
	}
}
