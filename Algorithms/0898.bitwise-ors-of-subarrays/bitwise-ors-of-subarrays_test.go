package problem0898

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	ans int
}{

	{
		[]int{0},
		1,
	},

	{
		[]int{1, 1, 2},
		3,
	},

	{
		[]int{1, 2, 4},
		6,
	},

	// 可以有多个 testcase
}

func Test_subarrayBitwiseORs(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, subarrayBitwiseORs(tc.A), "输入:%v", tc)
	}
}

func Benchmark_subarrayBitwiseORs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			subarrayBitwiseORs(tc.A)
		}
	}
}
