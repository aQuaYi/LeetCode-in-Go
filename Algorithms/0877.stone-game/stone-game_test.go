package problem0877

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	piles []int
	ans   bool
}{

	{
		[]int{1, 100, 2, 4},
		true,
	},

	{
		[]int{5, 3, 4, 5},
		true,
	},

	// 可以有多个 testcase
}

func Test_stoneGame(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, stoneGame(tc.piles), "输入:%v", tc)
	}
}

func Benchmark_stoneGame(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			stoneGame(tc.piles)
		}
	}
}
