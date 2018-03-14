package problem0303

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var nums = []int{-2, 0, 3, -5, 2, -1}

// tcs is testcase slice
var tcs = []struct {
	i, j int
	ans  int
}{

	{
		0, 2,
		1,
	},

	{
		2, 5,
		-1,
	},

	{
		0, 5,
		-3,
	},
	// 可以有多个 testcase
}

func Test_SumRange(t *testing.T) {
	ast := assert.New(t)

	na := Constructor(nums)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, na.SumRange(tc.i, tc.j), "输入:%v,%v,%v", nums, na.dp, tc)
	}
}

func Benchmark_Constructor(b *testing.B) {
	na := Constructor(nums)

	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			na.SumRange(tc.i, tc.j)
		}
	}
}
