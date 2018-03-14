package problem0242

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	t   string
	ans bool
}{

	{"anagram", "nagaram", true},
	{"rat", "car", false},
	{"at", "car", false},
	{"", "", true},
	{"世界和平", "和平世界", true},

	// 可以有多个 testcase
}

func Test_isAnagram(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, isAnagram(tc.s, tc.t), "输入:%v", tc)
	}
}

func Benchmark_isAnagram(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isAnagram(tc.s, tc.t)
		}
	}
}
