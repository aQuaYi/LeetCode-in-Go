package problem0899

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	S   string
	K   int
	ans string
}{

	{
		"cba",
		1,
		"acb",
	},

	{
		"edcba",
		5,
		"abcde",
	},

	{
		"dasfkasdlkfasdlkjflaksdjflksdajflkasdjflksdajflkasdjflkasdj",
		3,
		"aaaaaaaaadddddddddffffffffjjjjjjjkkkkkkkkkllllllllsssssssss",
	},

	{
		"baaca",
		3,
		"aaabc",
	},

	// 可以有多个 testcase
}

func Test_orderlyQueue(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, orderlyQueue(tc.S, tc.K), "输入:%v", tc)
	}
}

func Benchmark_orderlyQueue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			orderlyQueue(tc.S, tc.K)
		}
	}
}
