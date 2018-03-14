package problem0397

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n   int
	ans int
}{

	{8, 3},

	{1234567891, 42},

	{7, 4},

	// 可以有多个 testcase
}

func Test_integerReplacement(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, integerReplacement(tc.n), "输入:%v", tc)
	}
}

func Benchmark_integerReplacement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			integerReplacement(tc.n)
		}
	}
}
