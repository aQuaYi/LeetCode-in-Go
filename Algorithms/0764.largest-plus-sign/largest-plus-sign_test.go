package Problem0764

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	N int
 mines [][]int
	ans  int 
}{


	
	// 可以有多个 testcase
}

func Test_orderOfLargestPlusSign(t *testing.T) {
	ast := assert.New(t)
	
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, orderOfLargestPlusSign(tc.N, tc.mines), "输入:%v", tc)
	}
}

func Benchmark_orderOfLargestPlusSign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			orderOfLargestPlusSign(tc.N, tc.mines)
		}
	}
}
