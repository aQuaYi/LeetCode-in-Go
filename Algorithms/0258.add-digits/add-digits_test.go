package problem0258

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	num int
	ans int
}{

	{
		38,
		2,
	},

	// 可以有多个 testcase
}

func Test_addDigits(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, addDigits(tc.num), "输入:%v", tc)
	}
}

func Benchmark_addDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			addDigits(tc.num)
		}
	}
}
