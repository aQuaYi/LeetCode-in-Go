package problem0904

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	tree []int
	ans  int
}{

	{
		[]int{0},
		1,
	},

	{
		[]int{1, 2, 1},
		3,
	},

	{
		[]int{0, 1, 2, 2},
		3,
	},

	{
		[]int{1, 2, 3, 2, 2},
		4,
	},

	{
		[]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4},
		5,
	},

	// 可以有多个 testcase
}

func Test_totalFruit(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, totalFruit(tc.tree), "输入:%v", tc)
	}
}

func Benchmark_totalFruit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			totalFruit(tc.tree)
		}
	}
}
