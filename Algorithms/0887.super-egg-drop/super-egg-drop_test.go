package problem0887

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	K   int
	N   int
	ans int
}{

	{
		2,
		10000,
		141,
	},

	{
		5,
		10000,
		18,
	},

	{
		100,
		10000,
		14,
	},

	{
		1,
		2,
		2,
	},

	{
		2,
		6,
		3,
	},

	{
		3,
		14,
		4,
	},

	// 可以有多个 testcase
}

func Test_superEggDrop(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, superEggDrop(tc.K, tc.N), "输入:%v", tc)
	}
}

func Benchmark_superEggDrop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			superEggDrop(tc.K, tc.N)
		}
	}
}
