package problem0786

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	K   int
	ans []int
}{

	{
		[]int{1, 2, 3, 5},
		3,
		[]int{2, 5},
	},

	{
		[]int{1, 7},
		1,
		[]int{1, 7},
	},

	// 可以有多个 testcase
}

func Test_kthSmallestPrimeFraction(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, kthSmallestPrimeFraction(tc.A, tc.K), "输入:%v", tc)
	}
}

func Benchmark_kthSmallestPrimeFraction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			kthSmallestPrimeFraction(tc.A, tc.K)
		}
	}
}
