package problem0762

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A, B []int
	ans  []int
}{

	{
		[]int{47, 34, 51, 47, 47, 34},
		[]int{47, 51, 34, 34, 47, 47},
		[]int{5, 3, 1, 5, 5, 3},
	},

	{
		[]int{12, 28, 46, 32, 50, 12},
		[]int{50, 12, 32, 46, 28, 12},
		[]int{5, 4, 3, 2, 0, 5},
	},

	{
		[]int{12, 28, 46, 32, 50},
		[]int{50, 12, 32, 46, 28},
		[]int{1, 4, 3, 2, 0},
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, anagramMappings(tc.A, tc.B), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			anagramMappings(tc.A, tc.B)
		}
	}
}
