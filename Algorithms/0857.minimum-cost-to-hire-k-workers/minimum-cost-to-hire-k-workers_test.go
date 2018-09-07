package problem0857

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	quality []int
	wage    []int
	K       int
	ans     float64
}{

	{
		[]int{3000, 1, 20, 200, 1},
		[]int{30000, 8, 5, 50, 7},
		3,
		176,
	},

	{
		[]int{3000, 1, 20, 200, 1},
		[]int{4000, 8, 5, 50, 7},
		3,
		176,
	},

	{
		[]int{3000, 1, 10, 10, 1},
		[]int{4000, 8, 2, 2, 7},
		3,
		96,
	},

	{
		[]int{10, 20, 5},
		[]int{70, 50, 30},
		2,
		105.00000,
	},

	{
		[]int{3, 1, 10, 10, 1},
		[]int{4, 8, 2, 2, 7},
		3,
		30.66667,
	},

	// 可以有多个 testcase
}

func Test_mincostToHireWorkers(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		res := mincostToHireWorkers(tc.quality, tc.wage, tc.K)
		ast.InDelta(tc.ans, res, 0.00001, "输入:%v", tc)
	}
}

func Benchmark_mincostToHireWorkers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			mincostToHireWorkers(tc.quality, tc.wage, tc.K)
		}
	}
}
