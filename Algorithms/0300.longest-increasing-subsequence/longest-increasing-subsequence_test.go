package Problem0300

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  int
}{

	{
		[]int{10, 9, 2, 5, 3, 7, 101, 18},
		4,
	},

	// 可以有多个 testcase
}

func Test_lengthOfLIS(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, lengthOfLIS(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_lengthOfLIS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			lengthOfLIS(tc.nums)
		}
	}
}
