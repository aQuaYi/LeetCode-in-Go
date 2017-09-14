package Problem0115

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Problem0115(t *testing.T) {
	ast := assert.New(t)

	// tcs is testcase slice
	tcs := []struct {
		s   string
		t   string
		ans int
	}{

		{
			"rabbbit",
			"rabbit",
			3,
		},

		// 可以多个 testcase
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, numDistinct(tc.s, tc.t), "输入:%v", tc)
	}
}
