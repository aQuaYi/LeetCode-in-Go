package problem0335

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	x   []int
	ans bool
}{

	{
		[]int{1, 1, 2, 2, 3, 3, 4, 4, 10, 4, 4, 3, 3, 2, 2, 1, 1},
		false,
	},

	{
		[]int{1, 1, 4, 2, 2, 1, 1},
		false,
	},

	{
		[]int{1, 1, 3, 2, 1, 1, 1},
		true,
	},

	{
		[]int{1, 1, 2, 1, 1},
		true,
	},

	{
		[]int{2, 1, 3, 2, 2, 1},
		true,
	},

	{
		[]int{2, 1, 1, 2},
		true,
	},

	{
		[]int{1, 2, 2, 3, 4},
		false,
	},

	{
		[]int{1, 2, 3, 4},
		false,
	},

	{
		[]int{1, 1, 1, 1},
		true,
	},

	// 可以有多个 testcase
}

func Test_isSelfCrossing(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, isSelfCrossing(tc.x), "输入:%v", tc)
	}
}

func Benchmark_isSelfCrossing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isSelfCrossing(tc.x)
		}
	}
}
