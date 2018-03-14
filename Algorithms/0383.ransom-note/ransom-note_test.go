package problem0383

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	ransomNote string
	magazine   string
	ans        bool
}{

	{"a", "b", false},
	{"aa", "ab", false},
	{"aa", "aab", true},

	// 可以有多个 testcase
}

func Test_canConstruct(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, canConstruct(tc.ransomNote, tc.magazine), "输入:%v", tc)
	}
}

func Benchmark_canConstruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			canConstruct(tc.ransomNote, tc.magazine)
		}
	}
}
