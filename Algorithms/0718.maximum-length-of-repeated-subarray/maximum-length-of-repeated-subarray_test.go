package problem0718

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	a1, a2 []int
	ans    int
}{

	{
		[]int{1, 2, 3, 2, 1},
		[]int{3, 2, 4, 1, 7},
		2,
	},

	{
		[]int{1, 2, 3, 2, 1},
		[]int{3, 2, 1, 4, 7},
		3,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findLength(tc.a1, tc.a2), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findLength(tc.a1, tc.a2)
		}
	}
}
