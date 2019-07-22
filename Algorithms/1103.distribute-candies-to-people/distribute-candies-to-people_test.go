package problem1103

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	candies int
	num     int
	ans     []int
}{

	{
		7,
		4,
		[]int{1, 2, 3, 1},
	},

	{
		10,
		3,
		[]int{5, 2, 3},
	},

	// 可以有多个 testcase
}

func Test_distributeCandies(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, distributeCandies(tc.candies, tc.num), "输入:%v", tc)
	}
}

func Benchmark_distributeCandies(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			distributeCandies(tc.candies, tc.num)
		}
	}
}
