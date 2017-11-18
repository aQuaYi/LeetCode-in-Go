package Problem0546

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

		[]int{1, 3, 2, 2, 2, 3, 4, 3, 1},
		23,
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
