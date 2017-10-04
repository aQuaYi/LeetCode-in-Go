package Problem0221

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	matrix [][]byte
	ans    int
}{

	{
		[][]byte{
			[]byte("10100"),
			[]byte("10111"),
			[]byte("11111"),
			[]byte("10010"),
		},
		4,
	},

	// 可以有多个 testcase
}

func Test_maximalSquare(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, maximalSquare(tc.matrix), "输入:%v", tc)
	}
}

func Benchmark_maximalSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maximalSquare(tc.matrix)
		}
	}
}
