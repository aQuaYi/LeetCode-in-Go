package problem0916

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A, B []int
	ans  []string
}{

	// 可以有多个 testcase
}

func Test_wordSubsets(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, wordSubsets(tc.A, tc.B), "输入:%v", tc)
	}
}

func Benchmark_wordSubsets(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			wordSubsets(tc.A, tc.B)
		}
	}
}
