package Problem0279

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n int
	ans  int 
}{

	{
	12	,
		3,
	},

	// 可以有多个 testcase
}

func Test_numSquares(t *testing.T) {
	ast := assert.New(t)
	
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, numSquares(tc.n), "输入:%v", tc)
	}
}

func Benchmark_numSquares(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numSquares(tc.n)
		}
	}
}
