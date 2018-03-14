package problem0576

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	m, n, N, i, j int
	ans           int
}{

	{
		1,
		3,
		3,
		0,
		1,
		12,
	},

	{
		2,
		2,
		2,
		0,
		0,
		6,
	},

	{
		7,
		9,
		10,
		6,
		7,
		27982,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findPaths(tc.m, tc.n, tc.N, tc.i, tc.j), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findPaths(tc.m, tc.n, tc.N, tc.i, tc.j)
		}
	}
}
