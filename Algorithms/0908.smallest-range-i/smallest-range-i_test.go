package problem0908

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	K   int
	ans int
}{

	{
		[]int{1},
		0,
		0,
	},

	{
		[]int{0, 10},
		2,
		6,
	},

	{
		[]int{1, 0, 3, 6},
		3,
		0,
	},

	{
		[]int{1, 3, 6},
		3,
		0,
	},

	// 可以有多个 testcase
}

func Test_smallestRangeI(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, smallestRangeI(tc.A, tc.K), "输入:%v", tc)
	}
}

func Benchmark_smallestRangeI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			smallestRangeI(tc.A, tc.K)
		}
	}
}
