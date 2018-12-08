package problem0922

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A []int
}{

	{
		[]int{4, 2, 5, 7},
	},

	// 可以有多个 testcase
}

func Test_sortArrayByParityII(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.True(check(sortArrayByParityII(tc.A)))
	}
}

func Benchmark_sortArrayByParityII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			sortArrayByParityII(tc.A)
		}
	}
}

func check(a []int) bool {
	for i := range a {
		if a[i]%2 != i%2 {
			return false
		}
	}
	return true
}
