package problem0388

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	input string
	ans   int
}{

	{"dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext", 20},

	{"dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext\ndirasubdir2asubsubdir2aafile2.ext\n", 33},

	{"a\n\tb.c\n\tddddddddddddddddddddddddd", 5},

	{"dir\n\tsubdir1\n\tsubdir2", 0},

	{"dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext", 32},

	// 可以有多个 testcase
}

func Test_lengthLongestPath(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, lengthLongestPath(tc.input), "输入:%v", tc)
	}
}

func Benchmark_lengthLongestPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			lengthLongestPath(tc.input)
		}
	}
}
