package problem0227

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	ans int
}{

	{
		" 3/2 ",
		1,
	},

	{
		" 3+5 / 2 ",
		5,
	},

	{
		"3- 2*2",
		-1,
	},

	// 可以有多个 testcase
}

func Test_calculate(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, calculate(tc.s), "输入:%v", tc)
	}
}

func Benchmark_calculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			calculate(tc.s)
		}
	}
}
