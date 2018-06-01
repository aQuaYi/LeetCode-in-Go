package problem0836

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	rec1 []int
	rec2 []int
	ans  bool
}{

	{
		[]int{0, 0, 2, 2},
		[]int{1, 1, 3, 3},
		true,
	},

	{
		[]int{0, 0, 1, 1},
		[]int{1, 0, 2, 1},
		false,
	},

	// 可以有多个 testcase
}

func Test_isRectangleOverlap(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, isRectangleOverlap(tc.rec1, tc.rec2), "输入:%v", tc)
	}
}

func Benchmark_isRectangleOverlap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isRectangleOverlap(tc.rec1, tc.rec2)
		}
	}
}
