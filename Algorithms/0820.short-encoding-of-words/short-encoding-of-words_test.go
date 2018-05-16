package problem0820

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	words []string
	ans int
}{


	
	// 可以有多个 testcase
}

func Test_minimumLengthEncoding(t *testing.T) {
	ast := assert.New(t)
	
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, minimumLengthEncoding(tc.words), "输入:%v", tc)
	}
}

func Benchmark_minimumLengthEncoding(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minimumLengthEncoding(tc.words)
		}
	}
}
