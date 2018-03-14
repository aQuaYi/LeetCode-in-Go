package problem0455

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	g   []int
	s   []int
	ans int
}{

	{
		[]int{1, 2, 3},
		[]int{3},
		1,
	},

	{
		[]int{1, 2, 3},
		[]int{1, 1},
		1,
	},

	{
		[]int{1, 2},
		[]int{1, 2, 3},
		2,
	},

	// 可以有多个 testcase
}

func Test_findContentChildren(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findContentChildren(tc.g, tc.s), "输入:%v", tc)
	}
}

func Benchmark_findContentChildren(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findContentChildren(tc.g, tc.s)
		}
	}
}
