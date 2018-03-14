package problem0504

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	num int
	ans string
}{

	{
		0,
		"0",
	},

	{
		100,
		"202",
	},

	{
		-7,
		"-10",
	},

	// 可以有多个 testcase
}

func Test_convertToBase7(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, convertToBase7(tc.num), "输入:%v", tc)
	}
}

func Benchmark_convertToBase7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			convertToBase7(tc.num)
		}
	}
}
