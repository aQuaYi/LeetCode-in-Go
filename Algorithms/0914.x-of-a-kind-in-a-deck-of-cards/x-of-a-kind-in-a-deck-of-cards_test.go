package problem0914

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	deck []int
	ans  bool
}{

	{
		[]int{1, 2, 3, 4, 4, 3, 2, 1},
		true,
	},

	{
		[]int{1, 1, 1, 2, 2, 2, 3, 3},
		false,
	},

	{
		[]int{1},
		false,
	},

	{
		[]int{1, 1},
		true,
	},

	{
		[]int{1, 1, 2, 2, 2, 2},
		true,
	},

	// 可以有多个 testcase
}

func Test_myFunc(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, hasGroupsSizeX(tc.deck), "输入:%v", tc)
	}
}

func Benchmark_myFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			hasGroupsSizeX(tc.deck)
		}
	}
}
