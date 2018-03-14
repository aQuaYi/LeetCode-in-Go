package problem0678

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	ans bool
}{

	{
		"(())((())()()(*)(*()(())())())()()((()())((()))(*",
		false,
	},

	{
		")(",
		false,
	},

	{
		"()",
		true,
	},

	{
		"(*)",
		true,
	},

	{
		"(*))",
		true,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, checkValidString(tc.s), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			checkValidString(tc.s)
		}
	}
}
