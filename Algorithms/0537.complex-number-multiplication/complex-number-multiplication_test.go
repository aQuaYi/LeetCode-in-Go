package problem0537

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	a   string
	b   string
	ans string
}{

	{
		"1+1i",
		"1+1i",
		"0+2i",
	},

	{
		"1+-1i",
		"1+-1i",
		"0+-2i",
	},

	// 可以有多个 testcase
}

func Test_complexNumberMultiply(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, complexNumberMultiply(tc.a, tc.b), "输入:%v", tc)
	}
}

func Benchmark_complexNumberMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			complexNumberMultiply(tc.a, tc.b)
		}
	}
}
