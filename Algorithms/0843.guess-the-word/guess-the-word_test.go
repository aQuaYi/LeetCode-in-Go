package problem0843

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	this *Master
	ans Guess(word string) int
}{


	
	// 可以有多个 testcase
}

func Test_(t *testing.T) {
	ast := assert.New(t)
	
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, (tc.this), "输入:%v", tc)
	}
}

func Benchmark_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			(tc.this)
		}
	}
}
