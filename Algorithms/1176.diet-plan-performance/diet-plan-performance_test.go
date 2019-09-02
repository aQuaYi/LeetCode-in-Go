package problem1176

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	calories []int
	k        int
	lower    int
	upper    int
	ans      int
}{

	{
		[]int{1, 2, 3, 4, 5},
		1,
		3,
		3,
		0,
	},

	{
		[]int{3, 2},
		2,
		0,
		1,
		1,
	},

	{
		[]int{6, 5, 0, 0},
		2,
		1,
		5,
		0,
	},

	// 可以有多个 testcase
}

func Test_dietPlanPerformance(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, dietPlanPerformance(tc.calories, tc.k, tc.lower, tc.upper), "输入:%v", tc)
	}
}

func Benchmark_dietPlanPerformance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			dietPlanPerformance(tc.calories, tc.k, tc.lower, tc.upper)
		}
	}
}
