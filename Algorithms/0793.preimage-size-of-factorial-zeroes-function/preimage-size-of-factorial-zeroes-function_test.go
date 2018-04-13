package problem0793

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	K   int
	ans int
}{

	{
		0,
		5,
	},

	{
		17,
		0,
	},

	{
		11,
		0,
	},

	{
		79,
		0,
	},

	{
		5,
		0,
	},

	// 可以有多个 testcase
}

func Test_preimageSizeFZF(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, preimageSizeFZF(tc.K), "输入:%v", tc)
	}
}

func Benchmark_preimageSizeFZF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			preimageSizeFZF(tc.K)
		}
	}
}
