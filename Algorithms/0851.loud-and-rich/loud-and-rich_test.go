package problem0851

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	richer [][]int
	quiet  []int
	ans    []int
}{

	{
		[][]int{{1, 0}, {2, 1}, {3, 1}, {3, 7}, {4, 3}, {5, 3}, {6, 3}},
		[]int{3, 2, 5, 4, 6, 1, 7, 0},
		[]int{5, 5, 2, 5, 4, 5, 6, 7},
	},

	// 可以有多个 testcase
}

func Test_loudAndRich(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, loudAndRich(tc.richer, tc.quiet), "输入:%v", tc)
	}
}

func Benchmark_loudAndRich(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			loudAndRich(tc.richer, tc.quiet)
		}
	}
}
