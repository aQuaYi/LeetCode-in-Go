package problem0754

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	target int
	ans    int
}{

	{
		999999999,
		44721,
	},

	{
		-1000,
		47,
	},

	{
		1000,
		47,
	},

	{
		7,
		5,
	},

	{
		2,
		3,
	},

	{
		4,
		3,
	},

	{
		3,
		2,
	},

	{
		9,
		5,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, reachNumber(tc.target), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			reachNumber(tc.target)
		}
	}
}
