package problem0775

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	ans bool
}{

	{
		[]int{0, 1, 2},
		true,
	},

	{
		[]int{1, 0, 2},
		true,
	},

	{
		[]int{1, 2, 0},
		false,
	},

	// 可以有多个 testcase
}

func Test_isIdealPermutation(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, isIdealPermutation(tc.A), "输入:%v", tc)
	}
}

func Benchmark_isIdealPermutation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isIdealPermutation(tc.A)
		}
	}
}
