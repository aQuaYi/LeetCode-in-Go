package problem0241

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	input string
	ans   []int
}{

	{
		"2+2+2+2",
		[]int{8, 8, 8, 8, 8},
	},

	{
		"2-1-1",
		[]int{0, 2},
	},

	{
		"2*3-4*5",
		[]int{-34, -14, -10, -10, 10},
	},

	// 可以有多个 testcase
}

func Test_diffWaysToCompute(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ans := diffWaysToCompute(tc.input)
		sort.Ints(ans)
		ast.Equal(tc.ans, ans, "输入:%v", tc)
	}
}

func Benchmark_diffWaysToCompute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			diffWaysToCompute(tc.input)
		}
	}
}
