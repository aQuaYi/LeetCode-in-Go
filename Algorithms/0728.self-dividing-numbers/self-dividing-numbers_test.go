package problem0728

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	left, right int
	ans         []int
}{

	{
		1,
		22,
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22},
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, selfDividingNumbers(tc.left, tc.right), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			selfDividingNumbers(tc.left, tc.right)
		}
	}
}
