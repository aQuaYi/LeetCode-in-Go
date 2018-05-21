package problem0824

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	S   string
	ans string
}{

	{
		"I speak Goat Latin",
		"Imaa peaksmaaa oatGmaaaa atinLmaaaaa",
	},

	{
		"The quick brown fox jumped over the lazy dog",
		"heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa",
	},

	// 可以有多个 testcase
}

func Test_toGoatLatin(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, toGoatLatin(tc.S), "输入:%v", tc)
	}
}

func Benchmark_toGoatLatin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			toGoatLatin(tc.S)
		}
	}
}
