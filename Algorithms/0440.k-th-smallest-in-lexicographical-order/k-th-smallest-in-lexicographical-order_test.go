package problem0440

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n   int
	k   int
	ans int
}{

	{13, 2, 10},

	{987654321, 987654321, 99999999},

	{100, 10, 17},

	{1000000000, 1000000000, 999999999},

	// 可以有多个 testcase
}

func Test_findKthNumber(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findKthNumber(tc.n, tc.k), "输入:%v", tc)
	}
}

func Benchmark_findKthNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findKthNumber(tc.n, tc.k)
		}
	}
}
