package problem0591

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

	{
		"<DIV>>>  ![cdata[]] <![CDATA[<div>]>]>]>>]</DIV>",
		false,
	},

	{
		"<DIV>>>  ![cdata[]] <![CDATA[<div>]>]]>]]>>]</DIV>",
		true,
	},

	{
		"BBB <A>  <B> </A>   </B>",
		false,
	},

	{
		"<A>  <B> </A>   </B>",
		false,
	},

	{
		"<DIV>  div tag is not closed  <DIV>",
		false,
	},

	{
		"<DIV>  unmatched <  </DIV>",
		false,
	},

	{
		"<DIV> closed tags with invalid tag name  <b>123</b> </DIV>",
		false,
	},

	{
		"<DIV> unmatched tags with invalid tag name  </1234567890> and <CDATA[[]]>  </DIV>",
		false,
	},

	{
		"<DIV>  unmatched start tag <B>  and unmatched end tag </C>  </DIV>",
		false,
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
