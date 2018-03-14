package problem0598

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	m   int
	n   int
	ops [][]int
	ans int
}{

	{
		3,
		3,
		[][]int{{2, 2}, {3, 3}},
		4,
	},

	// 可以有多个 testcase
}

func Test_maxCount(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, maxCount(tc.m, tc.n, tc.ops), "输入:%v", tc)
	}
}

func Benchmark_maxCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxCount(tc.m, tc.n, tc.ops)
		}
	}
}
