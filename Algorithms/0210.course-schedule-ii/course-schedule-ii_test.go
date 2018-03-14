package problem0210

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	numCourses    int
	prerequisites [][]int
	ans           []int
}{

	{
		4,
		[][]int{
			[]int{1, 0},
			[]int{2, 1},
			[]int{2, 0},
			[]int{3, 0},
			[]int{3, 1},
			[]int{3, 2},
		},
		[]int{0, 1, 2, 3},
	},

	{
		4,
		[][]int{
			[]int{1, 0},
			[]int{2, 1},
			[]int{2, 0},
			[]int{3, 0},
			[]int{3, 1},
			[]int{3, 2},
			[]int{2, 3},
		},
		nil,
	},

	{
		2,
		[][]int{
			[]int{0, 1},
		},
		[]int{1, 0},
	},

	{
		2,
		[][]int{
			[]int{1, 0},
		},
		[]int{0, 1},
	},

	{
		2,
		[][]int{
			[]int{1, 0},
			[]int{0, 1},
		},
		nil,
	},
	// 可以有多个 testcase
}

func Test_findOrder(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findOrder(tc.numCourses, tc.prerequisites), "输入:%v", tc)
	}
}

func Benchmark_findOrder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findOrder(tc.numCourses, tc.prerequisites)
		}
	}
}
