package problem0407

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	heightMap [][]int
	ans       int
}{

	{
		[][]int{
			[]int{5, 5, 5, 1},
			[]int{5, 1, 1, 5},
		},
		0,
	},

	{
		[][]int{
			[]int{5, 5, 5, 1},
			[]int{5, 1, 1, 5},
			[]int{5, 1, 5, 5},
			[]int{5, 2, 5, 8},
		},
		3,
	},

	{
		[][]int{
			[]int{1, 4, 3, 1, 3, 2},
			[]int{3, 2, 1, 3, 2, 4},
			[]int{2, 3, 3, 2, 3, 1},
		},
		4,
	},

	{
		[][]int{
			[]int{12, 13, 1, 12},
			[]int{13, 4, 13, 12},
			[]int{13, 8, 10, 12},
			[]int{12, 13, 12, 12},
			[]int{13, 13, 13, 13},
		},
		14,
	},

	{
		[][]int{
			[]int{5, 5, 5, 5, 5},
			[]int{5, 1, 1, 1, 5},
			[]int{5, 1, 5, 5, 5},
			[]int{5, 1, 1, 1, 5},
			// 水会从这一行的 1 这里全部流出去，现有方法无法克服这种情况
			[]int{5, 5, 5, 1, 5},
		},
		0,
	},
	// 可以有多个 testcase
}

func Test_trapRainWater(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, trapRainWater(tc.heightMap), "输入:%v", tc)
	}
}

func Benchmark_trapRainWater(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			trapRainWater(tc.heightMap)
		}
	}
}
