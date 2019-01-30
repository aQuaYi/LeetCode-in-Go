package problem0932

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	N int
}{

	{
		4,
	},

	{
		5,
	},

	{
		10,
	},

	{
		1000,
	},

	// 可以有多个 testcase
}

func isBeautiful(A []int) bool {
	size := len(A)
	for i := 0; i < size; i++ {
		for j := i + 2; j < size; j++ {
			AiPlusAj := A[i] + A[j]
			for k := i + 1; k < j; k++ {
				Ak := A[k]
				if Ak*2 == AiPlusAj || Ak > size {
					fmt.Println(A[i], A[k], A[j])
					return false
				}
			}
		}
	}
	return true
}

func Test_checkArray(t *testing.T) {
	ast := assert.New(t)

	results := [][]int{
		{2, 1, 4, 5, 3},
		{3, 5, 1, 2, 4},
	}

	for _, r := range results {
		ast.True(isBeautiful(r))
	}
}

func Test_beautifulArray(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.True(isBeautiful(beautifulArray(tc.N)), "输入:%v", tc)
	}
}

func Benchmark_beautifulArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			beautifulArray(tc.N)
		}
	}
}
