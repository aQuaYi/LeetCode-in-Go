package problem0682

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	ops []string
	ans int
}{

	{
		[]string{"5", "2", "C", "D", "+"},
		30,
	},

	{
		[]string{"5", "-2", "4", "C", "D", "9", "+", "+"},
		27,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, calPoints(tc.ops), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			calPoints(tc.ops)
		}
	}
}
