package Problem0454

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	B   []int
	C   []int
	D   []int
	ans int
}{

	{
		[]int{1, 2},
		[]int{-2, -1},
		[]int{-1, 2},
		[]int{0, 2},
		2,
	},

	// 可以有多个 testcase
}

func Test_fourSumCount(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, fourSumCount(tc.A, tc.B, tc.C, tc.D), "输入:%v", tc)
	}
}

func Benchmark_fourSumCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			fourSumCount(tc.A, tc.B, tc.C, tc.D)
		}
	}
}
