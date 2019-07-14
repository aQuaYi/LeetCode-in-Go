package problem1108

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct{
	addr string
	ans string
}{
	{"1.1.1.1", "1[.]1[.]1[.]1"},
	{"255.100.50.0", "255[.]100[.]50[.]0"},
}

func Test_defangIPaddr(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, defangIPaddr(tc.addr), "输入：%v", tc)
	}
}

// Benchmark_defangIPaddr-8   	 5000000	       235 ns/op
func Benchmark_defangIPaddr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			defangIPaddr(tc.addr)
		}
	}
}

func Test_defangIPaddr2(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, defangIPaddr2(tc.addr), "输入：%v", tc)
	}
}

// Benchmark_defangIPaddr2-8   	 2000000	       709 ns/op
func Benchmark_defangIPaddr2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			defangIPaddr2(tc.addr)
		}
	}
}
