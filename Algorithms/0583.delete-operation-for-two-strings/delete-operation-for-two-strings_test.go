package problem0583

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	word1 string
	word2 string
	ans   int
}{

	{
		"bgswdnrsnzdrnmmyxpjmyqztaroernkakadxrzxhabbhlhexrdntmffrdhjjvqxhmfeowsi",
		"tegwdexzxrdvnziycqhdyswocagkthihuggxnaxvalbhhokbbceqmosxdhtjkgghharqnyl",
		96,
	},

	{
		"sea",
		"eat",
		2,
	},

	// 可以有多个 testcase
}

func Test_minDistance(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, minDistance(tc.word1, tc.word2), "输入:%v", tc)
	}
}

func Benchmark_minDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minDistance(tc.word1, tc.word2)
		}
	}
}
