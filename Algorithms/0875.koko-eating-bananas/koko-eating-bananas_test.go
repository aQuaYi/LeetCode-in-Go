package problem0875

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	piles []int
	H     int
	ans   int
}{

	{
		[]int{1, 2, 3, 4, 100},
		5,
		100,
	},

	{
		[]int{3, 6, 7, 10000},
		8,
		2000,
	},

	{
		[]int{3, 6, 7, 11},
		8,
		4,
	},

	{
		[]int{30, 11, 23, 4, 20},
		5,
		30,
	},

	{
		[]int{30, 11, 23, 4, 20},
		6,
		23,
	},

	// 可以有多个 testcase
}

func Test_minEatingSpeed(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, minEatingSpeed(tc.piles, tc.H), "输入:%v", tc)
	}
}

func Benchmark_minEatingSpeed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minEatingSpeed(tc.piles, tc.H)
		}
	}
}
