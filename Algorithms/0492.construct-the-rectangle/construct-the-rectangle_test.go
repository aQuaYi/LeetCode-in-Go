package problem0492

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	area int
	ans  []int
}{

	{
		10000000,
		[]int{3200, 3125},
	},

	{
		2,
		[]int{2, 1},
	},

	{
		4,
		[]int{2, 2},
	},

	// 可以有多个 testcase
}

func Test_constructRectangle(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, constructRectangle(tc.area), "输入:%v", tc)
	}
}

func Benchmark_constructRectangle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			constructRectangle(tc.area)
		}
	}
}
