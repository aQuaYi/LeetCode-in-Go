package problem0306

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	num string
	ans bool
}{

	{"0000000000", true},
	{"10112", true},
	{"101", true},
	{"112358", true},
	{"199100199", true},
	{"1991001991", false},
	{"11", false},
	{"00123", false},
	{"100100200", true},
	{"100000100000200000", true},
	{"9999991998", true},
	{"99911000", true},
	{"9981999", true},

	// 可以有多个 testcase
}

func Test_isAdditiveNumber(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, isAdditiveNumber(tc.num), "输入:%v", tc)
	}
}

func Benchmark_isAdditiveNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isAdditiveNumber(tc.num)
		}
	}
}
