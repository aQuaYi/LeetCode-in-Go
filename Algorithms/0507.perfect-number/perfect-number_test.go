package problem0507

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	num int
	ans bool
}{

	{
		100000000,
		false,
	},

	{
		1,
		false,
	},

	{
		132049,
		false,
	},

	{
		28,
		true,
	},

	// 可以有多个 testcase
}

func Test_checkPerfectNumber(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, checkPerfectNumber(tc.num), "输入:%v", tc)
	}
}

func Benchmark_checkPerfectNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			checkPerfectNumber(tc.num)
		}
	}
}
