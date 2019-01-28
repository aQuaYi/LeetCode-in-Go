package problem0191

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	num uint32
	ans int
}{

	{
		11, //0b00000000000000000000000000001011
		3,
	},

	{
		128, //0b00000000000000000000000010000000
		1,
	},

	{
		4294967293, //0b11111111111111111111111111111101
		31,
	},

	// 可以有多个 testcase
}

func Test_hammingWeight(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, hammingWeight(tc.num), "输入:%v", tc)
	}
}

func Benchmark_hammingWeight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			hammingWeight(tc.num)
		}
	}
}
