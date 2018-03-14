package problem0575

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	candies []int
	ans     int
}{

	{
		[]int{1, 1, 1, 1, 1, 1},
		1,
	},

	{
		[]int{1, 1, 2, 2, 3, 3},
		3,
	},

	{
		[]int{1, 1, 2, 3},
		2,
	},

	// 可以有多个 testcase
}

func Test_distributeCandies(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, distributeCandies(tc.candies), "输入:%v", tc)
	}
}

func Benchmark_distributeCandies(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			distributeCandies(tc.candies)
		}
	}
}
