package problem1051

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	heights []int
	ans     int
}{

	{
		[]int{1, 1, 4, 2, 1, 3},
		3,
	},

	// 可以有多个 testcase
}

func Test_heightChecker(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, heightChecker(tc.heights), "输入:%v", tc)
	}
}

func Benchmark_heightChecker(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			heightChecker(tc.heights)
		}
	}
}
