package problem0801

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	B   []int
	ans int
}{

	{
		[]int{1, 3, 5, 4},
		[]int{1, 2, 3, 7},
		1,
	},

	// 可以有多个 testcase
}

func Test_minSwap(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, minSwap(tc.A, tc.B), "输入:%v", tc)
	}
}

func Benchmark_minSwap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minSwap(tc.A, tc.B)
		}
	}
}
