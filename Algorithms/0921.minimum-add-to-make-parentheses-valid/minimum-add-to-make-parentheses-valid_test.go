package problem0921

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	S   string
	ans int
}{

	{
		"())",
		1,
	},

	{
		"(((",
		3,
	},

	{
		"()",
		0,
	},

	{
		"()))((",
		4,
	},

	// 可以有多个 testcase
}

func Test_minAddToMakeValid(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, minAddToMakeValid(tc.S), "输入:%v", tc)
	}
}

func Benchmark_minAddToMakeValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minAddToMakeValid(tc.S)
		}
	}
}
