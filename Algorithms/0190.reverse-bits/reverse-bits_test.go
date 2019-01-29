package problem0190

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	num uint32
	ans uint32
}{

	{
		43261596,  // 0b00000010100101000001111010011100
		964176192, // 0b00111001011110000010100101000000
	},

	{
		4294967293, // 0b11111111111111111111111111111101
		3221225471, // 0b10101111110010110010011101101001
	},

	// 可以有多个 testcase
}

func Test_reverseBits(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, reverseBits(tc.num), "输入:%v", tc)
	}
}

func Benchmark_reverseBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			reverseBits(tc.num)
		}
	}
}
