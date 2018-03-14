package problem0546

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	boxes []int
	ans   int
}{

	{
		[]int{},
		0,
	},

	{
		[]int{1, 3, 2, 2, 2, 3, 4, 3, 1},
		23,
	},

	{
		[]int{3, 8, 8, 5, 5, 3, 9, 2, 4, 4, 6, 5, 8, 4, 8, 6, 9, 6, 2, 8, 6, 4, 1, 9, 5, 3, 10, 5, 3, 3, 9, 8, 8, 6, 5, 3, 7, 4, 9, 6, 3, 9, 4, 3, 5, 10, 7, 6, 10, 7},
		136,
	},

	// 可以有多个 testcase
}

func Test_removeBoxes(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, removeBoxes(tc.boxes), "输入:%v", tc)
	}
}

func Benchmark_removeBoxes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			removeBoxes(tc.boxes)
		}
	}
}
