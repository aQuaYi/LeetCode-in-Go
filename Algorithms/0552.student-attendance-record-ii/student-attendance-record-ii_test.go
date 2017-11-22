package Problem0552

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

	{3, 19},

	{2, 8},

	{4, 43},

	{5, 94},

	{10, 3536},

	// {100, 985598218},

	// {100000, 749184020},

	// 可以有多个 testcase
}

func Test_checkRecord(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, checkRecord(tc.n), "输入:%v", tc)
	}
}

func Test_afterL(t *testing.T) {
	ast := assert.New(t)

	actual := afterL(2)
	expected := 7
	ast.Equal(expected, actual)
}

func Benchmark_checkRecord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			checkRecord(tc.n)
		}
	}
}
