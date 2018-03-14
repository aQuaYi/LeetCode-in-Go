package problem0475

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	houses  []int
	heaters []int
	ans     int
}{

	{
		[]int{1, 2, 3, 5, 15},
		[]int{2, 30},
		13,
	},

	{
		[]int{},
		[]int{3},
		0,
	},

	{
		[]int{3},
		[]int{3},
		0,
	},

	{
		[]int{1, 2, 3},
		[]int{2},
		1,
	},

	{
		[]int{3, 2, 4},
		[]int{3},
		1,
	},

	{
		[]int{3, 5, 7, 9},
		[]int{8, 4},
		1,
	},

	{
		[]int{3, 5, 6, 7, 9},
		[]int{8, 4},
		2,
	},

	{
		[]int{1, 2, 3, 4},
		[]int{1, 4},
		1,
	},

	// 可以有多个 testcase
}

func Test_findRadius(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findRadius(tc.houses, tc.heaters), "输入:%v", tc)
	}
}

func Benchmark_findRadius(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findRadius(tc.houses, tc.heaters)
		}
	}
}
