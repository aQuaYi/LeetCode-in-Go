package problem0557

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

	{
		"Let's take LeetCode contest",
		"s'teL ekat edoCteeL tsetnoc",
	},

	// 可以有多个 testcase
}

func Test_reverseWords(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, reverseWords(tc.s), "输入:%v", tc)
	}
}

func Benchmark_reverseWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			reverseWords(tc.s)
		}
	}
}
