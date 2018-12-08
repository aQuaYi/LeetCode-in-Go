package problem0922

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	ans []int
}{

	{
		[]int{4, 2, 5, 7},
		[]int{4, 5, 2, 7},
	},

	// 可以有多个 testcase
}

func Test_sortArrayByParityII(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, sortArrayByParityII(tc.A), "输入:%v", tc)
	}
}

func Benchmark_sortArrayByParityII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			sortArrayByParityII(tc.A)
		}
	}
}
