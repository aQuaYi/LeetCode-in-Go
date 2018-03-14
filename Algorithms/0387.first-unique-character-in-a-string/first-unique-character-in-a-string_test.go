package problem0387

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	ans int
}{

	{
		"leetcode",
		0,
	},

	{
		"loveleetcode",
		2,
	},

	{"aabbcc", -1},

	// 可以有多个 testcase
}

func Test_firstUniqChar(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, firstUniqChar(tc.s), "输入:%v", tc)
	}
}

func Benchmark_firstUniqChar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			firstUniqChar(tc.s)
		}
	}
}
