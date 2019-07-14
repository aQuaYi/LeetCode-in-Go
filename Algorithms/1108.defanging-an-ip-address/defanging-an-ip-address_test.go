package problem1108

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	address string
	ans     string
}{

	{
		"1.1.1.1",
		"1[.]1[.]1[.]1",
	},

	{
		"255.100.50.0",
		"255[.]100[.]50[.]0",
	},

	// 可以有多个 testcase
}

func Test_defangIPaddr(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, defangIPaddr(tc.address), "输入:%v", tc)
	}
}

func Benchmark_defangIPaddr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			defangIPaddr(tc.address)
		}
	}
}
