package problem0412

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n   int
	ans []string
}{

	{
		15,
		[]string{
			"1",
			"2",
			"Fizz",
			"4",
			"Buzz",
			"Fizz",
			"7",
			"8",
			"Fizz",
			"Buzz",
			"11",
			"Fizz",
			"13",
			"14",
			"FizzBuzz",
		},
	},

	// 可以有多个 testcase
}

func Test_fizzBuzz(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, fizzBuzz(tc.n), "输入:%v", tc)
	}
}

func Benchmark_fizzBuzz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			fizzBuzz(tc.n)
		}
	}
}
