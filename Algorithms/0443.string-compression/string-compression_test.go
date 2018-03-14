package problem0443

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	chars      []byte
	ans        int
	ansInPlace []byte
}{

	{
		[]byte("abaa2"),
		5,
		[]byte("aba22"),
	},

	{
		[]byte("a"),
		1,
		[]byte("a"),
	},

	{
		[]byte("abbbbbbbbbbbb"),
		4,
		[]byte("ab12"),
	},

	{
		[]byte("aabbccc"),
		6,
		[]byte("a2b2c3"),
	},

	// 可以有多个 testcase
}

func Test_compress(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, compress(tc.chars))

		// 检查 inPlace 修改的情况
		ast.Equal(string(tc.ansInPlace), string(tc.chars)[:len(tc.ansInPlace)])
	}
}

func Benchmark_compress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			compress(tc.chars)
		}
	}
}
