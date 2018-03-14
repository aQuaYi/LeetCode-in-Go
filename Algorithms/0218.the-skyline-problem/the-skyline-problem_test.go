package problem0218

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	buildings [][]int
	ans       [][]int
}{

	{
		[][]int{
			[]int{2, 9, 10},
			[]int{3, 7, 15},
			[]int{5, 12, 12},
			[]int{15, 20, 10},
			[]int{19, 24, 8},
		},
		[][]int{
			[]int{2, 10},
			[]int{3, 15},
			[]int{7, 12},
			[]int{12, 0},
			[]int{15, 10},
			[]int{20, 8},
			[]int{24, 0},
		},
	},

	{
		[][]int{
			[]int{1, 6, 1},
			[]int{2, 3, 2},
			[]int{4, 5, 2},
		},
		[][]int{
			[]int{1, 1},
			[]int{2, 2},
			[]int{3, 1},
			[]int{4, 2},
			[]int{5, 1},
			[]int{6, 0},
		},
	},

	{
		[][]int{
			[]int{1, 2, 1},
			[]int{2, 3, 2},
		},
		[][]int{
			[]int{1, 1},
			[]int{2, 2},
			[]int{3, 0},
		},
	},

	{
		[][]int{
			[]int{1, 2, 1},
			[]int{1, 2, 2},
			[]int{1, 2, 3},
		},
		[][]int{
			[]int{1, 3},
			[]int{2, 0},
		},
	},
	// 可以有多个 testcase
}

func Test_getSkyline(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, getSkyline(tc.buildings), "输入:%v", tc)
	}
}

func Benchmark_getSkyline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			getSkyline(tc.buildings)
		}
	}
}
