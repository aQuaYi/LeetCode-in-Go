package problem0648

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	dict     []string
	sentence string
	ans      string
}{

	{
		[]string{"cat", "bat", "rat"},
		"the cattle was rattled by the battery",
		"the cat was rat by the bat",
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, replaceWords(tc.dict, tc.sentence), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			replaceWords(tc.dict, tc.sentence)
		}
	}
}
