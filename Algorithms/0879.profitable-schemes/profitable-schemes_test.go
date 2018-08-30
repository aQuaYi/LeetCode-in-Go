package problem0879

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	G      int
	P      int
	group  []int
	profit []int
	ans    int
}{

	{
		10,
		5,
		[]int{2, 3, 5},
		[]int{6, 7, 8},
		7,
	},

	{
		5,
		3,
		[]int{2, 2},
		[]int{2, 3},
		2,
	},

	{
		5,
		0,
		[]int{2, 2},
		[]int{2, 3},
		4,
	},

	// 可以有多个 testcase
}

func Test_profitableSchemes(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, profitableSchemes(tc.G, tc.P, tc.group, tc.profit), "输入:%v", tc)
	}
}

func Benchmark_profitableSchemes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			profitableSchemes(tc.G, tc.P, tc.group, tc.profit)
		}
	}
}
