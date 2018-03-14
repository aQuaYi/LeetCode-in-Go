package problem0390

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

	{9, 6},

	{123456789, 56614230},

	// 可以有多个 testcase
}

func Test_lastRemaining(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, lastRemaining(tc.n), "输入:%v", tc)
	}
}

func Benchmark_lastRemaining(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			lastRemaining(tc.n)
		}
	}
}
