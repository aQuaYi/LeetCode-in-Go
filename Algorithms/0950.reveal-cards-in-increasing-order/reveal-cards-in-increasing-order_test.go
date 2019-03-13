package problem0950

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	deck []int
	ans  []int
}{

	{
		[]int{17, 13, 11, 2, 3, 5, 7},
		[]int{2, 13, 3, 11, 5, 17, 7},
	},

	// 可以有多个 testcase
}

func Test_deckRevealedIncreasing(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, deckRevealedIncreasing(tc.deck), "输入:%v", tc)
	}
}

func Benchmark_deckRevealedIncreasing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			deckRevealedIncreasing(tc.deck)
		}
	}
}
