package Problem0502

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	k       int
	W       int
	Profits []int
	Capital []int
	ans     int
}{

	{
		2,
		1,
		[]int{1, 2, 3},
		[]int{1, 1, 1},
		6,
	},

	{
		2,
		0,
		[]int{1, 2, 3},
		[]int{0, 1, 1},
		4,
	},

	// 可以有多个 testcase
}

func Test_findMaximizedCapital(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findMaximizedCapital(tc.k, tc.W, tc.Profits, tc.Capital), "输入:%v", tc)
	}
}

func Benchmark_findMaximizedCapital(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findMaximizedCapital(tc.k, tc.W, tc.Profits, tc.Capital)
		}
	}
}
