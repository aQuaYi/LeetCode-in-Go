package Problem0468

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	IP string
	ans  string 
}{


	
	// 可以有多个 testcase
}

func Test_validIPAddress(t *testing.T) {
	ast := assert.New(t)
	
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, validIPAddress(tc.IP), "输入:%v", tc)
	}
}

func Benchmark_validIPAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			validIPAddress(tc.IP)
		}
	}
}
