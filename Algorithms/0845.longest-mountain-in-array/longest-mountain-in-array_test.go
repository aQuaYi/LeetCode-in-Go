package problem0845

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	ans int
}{

	{
		[]int{2, 3, 3, 2, 0, 2},
		0,
	},

	{
		[]int{1, 2, 0, 2, 0, 2},
		3,
	},

	{
		[]int{2, 1, 4, 7, 3, 2, 5},
		5,
	},

	{
		[]int{2, 2, 2},
		0,
	},

	{
		[]int{2, 3},
		0,
	},

	// 可以有多个 testcase
}

func Test_longestMountain(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, longestMountain(tc.A), "输入:%v", tc)
	}
}

func Benchmark_longestMountain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			longestMountain(tc.A)
		}
	}
}
