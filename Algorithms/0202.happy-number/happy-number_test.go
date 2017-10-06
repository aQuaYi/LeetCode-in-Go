package Problem0202

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n   int
	ans bool
}{

	{
		77,
		false,
	},

	{
		19,
		true,
	},

	// 可以有多个 testcase
}

func Test_isHappy(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, isHappy(tc.n), "输入:%v", tc)
	}
}

func Benchmark_isHappy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isHappy(tc.n)
		}
	}
}
