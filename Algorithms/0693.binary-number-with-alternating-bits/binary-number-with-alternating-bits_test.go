package problem0693

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
		5,
		true,
	},

	{
		3,
		false,
	},

	{
		213,
		false,
	},

	{
		7,
		false,
	},

	{
		11,
		false,
	},

	{
		10,
		true,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, hasAlternatingBits(tc.n), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			hasAlternatingBits(tc.n)
		}
	}
}
