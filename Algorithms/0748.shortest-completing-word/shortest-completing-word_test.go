package problem0748

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	licensePlate string
	words        []string
	ans          string
}{

	{
		"1s3 PSt",
		[]string{"step", "steps", "stripe", "stepple"},
		"steps",
	},

	{
		"1s3 456",
		[]string{"looks", "pest", "stew", "show"},
		"pest",
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, shortestCompletingWord(tc.licensePlate, tc.words), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			shortestCompletingWord(tc.licensePlate, tc.words)
		}
	}
}
