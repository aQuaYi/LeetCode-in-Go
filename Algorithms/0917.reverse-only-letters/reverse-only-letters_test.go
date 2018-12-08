package problem0917

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
		"ab-cd",
		"dc-ba",
	},

	{
		"a-bC-dEf-ghIj",
		"j-Ih-gfE-dCba",
	},

	{
		"Test1ng-Leet=code-Q!",
		"Qedo1ct-eeLg=ntse-T!",
	},

	// 可以有多个 testcase
}

func Test_reverseOnlyLetters(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, reverseOnlyLetters(tc.S), "输入:%v", tc)
	}
}

func Benchmark_reverseOnlyLetters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			reverseOnlyLetters(tc.S)
		}
	}
}
