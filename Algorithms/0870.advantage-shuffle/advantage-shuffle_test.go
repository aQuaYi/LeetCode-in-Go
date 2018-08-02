package problem0870

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	B   []int
	ans []int
}{

	{
		[]int{15448, 14234, 13574, 19893, 6475},
		[]int{14234, 6475, 19893, 15448, 13574},
		[]int{15448, 13574, 6475, 19893, 14234},
	},

	{
		[]int{2, 7, 11, 15},
		[]int{1, 10, 4, 11},
		[]int{2, 11, 7, 15},
	},

	{
		[]int{12, 24, 8, 32},
		[]int{13, 25, 32, 11},
		[]int{24, 32, 8, 12},
	},

	// 可以有多个 testcase
}

func Test_advantageCount(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, advantageCount(tc.A, tc.B), "输入:%v", tc)
	}
}

func Benchmark_advantageCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			advantageCount(tc.A, tc.B)
		}
	}
}
