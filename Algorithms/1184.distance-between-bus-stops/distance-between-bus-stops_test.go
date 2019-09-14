package problem1184

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	distance    []int
	start       int
	destination int
	ans         int
}{

	{
		[]int{7, 10, 1, 12, 11, 14, 5, 0},
		7,
		2,
		17,
	},

	{
		[]int{1, 2, 3, 4},
		0,
		1,
		1,
	},

	{
		[]int{1, 2, 3, 4},
		0,
		2,
		3,
	},

	{
		[]int{1, 2, 3, 4},
		0,
		3,
		4,
	},

	// 可以有多个 testcase
}

func Test_distanceBetweenBusStops(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, distanceBetweenBusStops(tc.distance, tc.start, tc.destination), "输入:%v", tc)
	}
}

func Benchmark_distanceBetweenBusStops(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			distanceBetweenBusStops(tc.distance, tc.start, tc.destination)
		}
	}
}
