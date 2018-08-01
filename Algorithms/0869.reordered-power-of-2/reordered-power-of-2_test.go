package problem0869

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	N   int
	ans bool
}{

	{
		1,
		true,
	},

	{
		10,
		false,
	},

	{
		16,
		true,
	},

	{
		24,
		false,
	},

	{
		46,
		true,
	},

	// 可以有多个 testcase
}

func Test_reorderedPowerOf2(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, reorderedPowerOf2(tc.N), "输入:%v", tc)
	}
}

func Benchmark_reorderedPowerOf2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			reorderedPowerOf2(tc.N)
		}
	}
}
