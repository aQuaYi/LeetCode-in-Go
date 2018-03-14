package problem0712

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s1, s2 string
	ans    int
}{

	{
		"delete",
		"leet",
		403,
	},

	{
		"sea",
		"eat",
		231,
	},

	// 可以有多个 testcase
}

func Test_minimumDeleteSum(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, minimumDeleteSum(tc.s1, tc.s2), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minimumDeleteSum(tc.s1, tc.s2)
		}
	}
}
