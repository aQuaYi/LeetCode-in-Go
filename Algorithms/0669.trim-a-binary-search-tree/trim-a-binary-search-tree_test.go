package Problem0669

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	pre, in []int
	L, R    int
	ansPost []int
}{

	{
		[]int{1, 0, 2},
		[]int{0, 1, 2},
		1,
		2,
		[]int{2, 1},
	},

	{
		[]int{3, 0, 2, 1, 4},
		[]int{0, 1, 2, 3, 4},
		1,
		3,
		[]int{3, 2, 1},
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, trimBST(tc.para), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			trimBST(tc.para)
		}
	}
}
