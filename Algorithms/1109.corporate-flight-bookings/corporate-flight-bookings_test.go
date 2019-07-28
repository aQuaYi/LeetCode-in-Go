package problem1109

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	bookings [][]int
	n        int
	ans      []int
}{

	{
		[][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}},
		5,
		[]int{10, 55, 45, 25, 25},
	},

	// 可以有多个 testcase
}

func Test_corpFlightBookings(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, corpFlightBookings(tc.bookings, tc.n), "输入:%v", tc)
	}
}

func Benchmark_corpFlightBookings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			corpFlightBookings(tc.bookings, tc.n)
		}
	}
}
