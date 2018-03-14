package problem0344

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	ans string
}{

	{"hello", "olleh"},
	{"world", "dlrow"},

	// 可以有多个 testcase
}

func Test_reverseString(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, reverseString(tc.s), "输入:%v", tc)
	}
}

func Benchmark_reverseString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			reverseString(tc.s)
		}
	}
}
