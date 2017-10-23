package Problem0363

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	matrix [][]int
	k      int
	ans    int
}{

	{
		[][]int{
			[]int{1, 0, 1},
			[]int{0, -2, 3},
		},
		2,
		2,
	},

	// 可以有多个 testcase
}

func Test_maxSumSubmatrix(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, maxSumSubmatrix(tc.matrix, tc.k), "输入:%v", tc)
	}
}

func Benchmark_maxSumSubmatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxSumSubmatrix(tc.matrix, tc.k)
		}
	}
}
