package problem0831

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	S   string
	ans string
}{

	{
		"LeetCode@LeetCode.com",
		"l*****e@leetcode.com",
	},

	{
		"AB@qq.com",
		"a*****b@qq.com",
	},

	{
		"1(234)567-890",
		"***-***-7890",
	},

	{
		"86-(10)12345678",
		"+**-***-***-5678",
	},

	// 可以有多个 testcase
}

func Test_maskPII(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, maskPII(tc.S), "输入:%v", tc)
	}
}

func Benchmark_maskPII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maskPII(tc.S)
		}
	}
}
