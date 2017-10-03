package Problem0218

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	buildings [][]int
	ans  [][]int 
}{

	{
		,
		,
	},

	// 可以有多个 testcase
}

func Test_getSkyline(t *testing.T) {
	ast := assert.New(t)
	
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, getSkyline(tc.buildings), "输入:%v", tc)
	}
}

func Benchmark_getSkyline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			getSkyline(tc.buildings)
		}
	}
}
