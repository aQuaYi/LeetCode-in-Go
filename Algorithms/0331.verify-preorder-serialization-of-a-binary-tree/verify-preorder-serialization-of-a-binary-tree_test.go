package Problem0331

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	preorder string
	ans      bool
}{

	{
		"9,3,4,#,#,1,#,#,2,#,6,#,#",
		true,
	},

	{
		"1,#",
		false,
	},

	{
		"#",
		true,
	},

	{
		"9,#,#,1",
		false,
	},

	// 可以有多个 testcase
}

func Test_isValidSerialization(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, isValidSerialization(tc.preorder), "输入:%v", tc)
	}
}

func Benchmark_isValidSerialization(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isValidSerialization(tc.preorder)
		}
	}
}
