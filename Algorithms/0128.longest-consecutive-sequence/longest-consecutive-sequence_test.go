package Problem0128

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
		[]int{0, -1},
		2,
	},

	{
		[]int{},
		0,
	},

	{
		[]int{0},
		1,
	},

	{
		[]int{100, 4, 200, 1, 3, 2},
		4,
	},

	// 可以有多个 testcase
}

func Test_longestConsecutive(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, longestConsecutive(tc.nums), "输入:%!s(MISSING)", tc)
	}
}

func Benchmark_longestConsecutive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			longestConsecutive(tc.nums)
		}
	}
}
