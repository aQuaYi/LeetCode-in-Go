package Problem0443

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	chars []byte
	ans   int
}{

	{
		[]byte("abaa2"),
		5,
	},

	{
		[]byte("a"),
		1,
	},

	{
		[]byte("abbbbbbbbbbbb"),
		4,
	},

	{
		[]byte("aabbccc"),
		6,
	},

	// 可以有多个 testcase
}

func Test_compress(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, compress(tc.chars), "输入:%v", tc)
	}
}

func Benchmark_compress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			compress(tc.chars)
		}
	}
}
