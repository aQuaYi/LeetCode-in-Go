package Problem0591

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	code string
	ans  bool
}{

	{
		"<DIV>This is the first line <![CDATA[<div>]]></DIV>",
		true,
	},

	// 可以有多个 testcase
}

func Test_isValid(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, isValid(tc.code), "输入:%v", tc)
	}
}

func Benchmark_isValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isValid(tc.code)
		}
	}
}
