package problem1095

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	target      int
	mountainArr []int
	ans         int
}{

	{
		2,
		[]int{1, 5, 2},
		2,
	},

	{
		2,
		[]int{1, 2, 3, 4, 5, 3, 1},
		1,
	},

	{
		3,
		[]int{1, 2, 3, 4, 5, 3, 1},
		2,
	},

	{
		3,
		[]int{0, 1, 2, 4, 2, 1},
		-1,
	},

	// 可以有多个 testcase
}

func Test_findInMountainArray(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		m := MountainArray(tc.mountainArr)
		ast.Equal(tc.ans, findInMountainArray(tc.target, &m), "输入:%v", tc)
	}
}

func Benchmark_findInMountainArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			m := MountainArray(tc.mountainArr)
			findInMountainArray(tc.target, &m)
		}
	}
}
