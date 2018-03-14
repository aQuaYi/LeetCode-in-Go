package problem0125

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Problem0125(t *testing.T) {
	ast := assert.New(t)

	// tcs is testcase slice
	tcs := []struct {
		s   string
		ans bool
	}{

		{
			"0p",
			false,
		},

		{
			"0",
			true,
		},

		{
			"race a car",
			false,
		},

		{
			"A man, a plan, a canal: Panama",
			true,
		},

		// 可以多个 testcase
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, isPalindrome(tc.s), "输入:%v", tc)
	}
}
