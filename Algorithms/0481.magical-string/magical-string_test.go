package problem0481

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n   int
	ans int
}{

	// {100000, 49972},

	// {10000, 4996},

	// {1000, 502},

	// {100, 49},

	{10, 5},

	{1, 1},

	// 可以有多个 testcase
}

func Test_magicalString(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, magicalString(tc.n), "输入:%v", tc)
	}
}

func Benchmark_magicalString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			magicalString(tc.n)
		}
	}
}
