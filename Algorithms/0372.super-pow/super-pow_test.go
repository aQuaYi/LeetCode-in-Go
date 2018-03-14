package problem0372

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	a   int
	b   []int
	ans int
}{

	{
		2,
		[]int{3, 4, 5},
		694,
	},

	{
		2,
		[]int{3},
		8,
	},

	{
		2,
		[]int{1, 0},
		1024,
	},

	// 可以有多个 testcase
}

func Test_superPow(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, superPow(tc.a, tc.b), "输入:%v", tc)
	}
}

func Benchmark_superPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			superPow(tc.a, tc.b)
		}
	}
}
