package Problem0315

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  []int
}{

	{
		[]int{5, 2, 6, 1},
		[]int{2, 1, 1, 0},
	},

	// 可以有多个 testcase
}

func Test_countSmaller(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, countSmaller(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_countSmaller(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			countSmaller(tc.nums)
		}
	}
}
