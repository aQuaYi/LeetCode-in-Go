package Problem0472

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	words []string
	ans   []string
}{

	{
		[]string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"},
		[]string{"catsdogcats", "dogcatsdog", "ratcatdogcat"},
	},

	// 可以有多个 testcase
}

func Test_findAllConcatenatedWordsInADict(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findAllConcatenatedWordsInADict(tc.words), "输入:%v", tc)
	}
}

func Benchmark_findAllConcatenatedWordsInADict(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findAllConcatenatedWordsInADict(tc.words)
		}
	}
}
