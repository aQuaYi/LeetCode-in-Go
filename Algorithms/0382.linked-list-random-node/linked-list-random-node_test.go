package Problem0382

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	head *ListNode
	ans  Solution 
}{

	{ },
	{ },

	// 可以有多个 testcase
}

func Test_Constructor(t *testing.T) {
	ast := assert.New(t)
	
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, Constructor(tc.head), "输入:%v", tc)
	}
}

func Benchmark_Constructor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			Constructor(tc.head)
		}
	}
}
