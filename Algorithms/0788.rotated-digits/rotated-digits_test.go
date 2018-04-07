package problem0788

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	N   int
	ans int
}{

	{
		9876,
		2257,
	},

	{
		10000,
		2320,
	},

	{
		100,
		40,
	},

	{
		10,
		4,
	},

	// 可以有多个 testcase
}

func Test_rotatedDigits(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, rotatedDigits(tc.N), "输入:%v", tc)
	}
}

func Benchmark_rotatedDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			rotatedDigits(tc.N)
		}
	}
}
