package problem0798

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
		[]int{2, 3, 1, 4, 0},
		3,
	},

	{
		[]int{1, 3, 0, 2, 4},
		0,
	},

	// 可以有多个 testcase
}

func Test_bestRotation(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, bestRotation(tc.A), "输入:%v", tc)
	}
}

func Benchmark_bestRotation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			bestRotation(tc.A)
		}
	}
}
