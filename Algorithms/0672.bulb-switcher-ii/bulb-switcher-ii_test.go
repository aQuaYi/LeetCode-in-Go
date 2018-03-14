package problem0672

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n, m int
	ans  int
}{

	{
		1,
		1,
		2,
	},

	{
		2,
		1,
		3,
	},

	{
		2,
		4,
		4,
	},

	{
		3,
		4,
		8,
	},

	{
		3,
		0,
		1,
	},

	{
		3,
		3,
		8,
	},

	{
		3,
		2,
		7,
	},

	{
		3,
		1,
		4,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, flipLights(tc.n, tc.m), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			flipLights(tc.n, tc.m)
		}
	}
}
