package problem0779

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	N   int
	K   int
	ans int
}{

	{
		30,
		434991989,
		0,
	},

	{
		1,
		1,
		0,
	},

	{
		2,
		1,
		0,
	},

	{
		2,
		2,
		1,
	},

	{
		4,
		5,
		1,
	},

	// 可以有多个 testcase
}

func Test_kthGrammar(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, kthGrammar(tc.N, tc.K), "输入:%v", tc)
	}
}

func Benchmark_kthGrammar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			kthGrammar(tc.N, tc.K)
		}
	}
}
