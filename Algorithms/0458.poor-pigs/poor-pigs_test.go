package problem0458

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	buckets       int
	minutesToDie  int
	minutesToTest int
	ans           int
}{

	{
		1,
		1,
		1,
		0,
	},

	{
		1000,
		12,
		60,
		4,
	},

	{
		1000,
		15,
		60,
		5,
	},

	// 可以有多个 testcase
}

func Test_myFunc(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, poorPigs(tc.buckets, tc.minutesToDie, tc.minutesToTest), "输入:%v", tc)
	}
}

func Benchmark_myFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			poorPigs(tc.buckets, tc.minutesToDie, tc.minutesToTest)
		}
	}
}
