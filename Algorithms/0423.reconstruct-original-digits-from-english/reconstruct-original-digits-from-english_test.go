package problem0423

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	ans string
}{

	{
		"",
		"",
	},

	{
		"owoztneoer",
		"012",
	},

	{
		"zeroonetwothreefourfivesixseveneightnien",
		"0123456789",
	},

	{
		"fviefuro",
		"45",
	},

	// 可以有多个 testcase
}

func Test_originalDigits(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, originalDigits(tc.s), "输入:%v", tc)
	}
}

func Benchmark_originalDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		originalDigits("zeroonetwothreefourfivesixseveneightnienzeroonetwothreefourfivesixseveneightnienzeroonetwothreefourfivesixseveneightnienzeroonetwothreefourfivesixseveneightnienzeroonetwothreefourfivesixseveneightnienzeroonetwothreefourfivesixseveneightnienzeroonetwothreefourfivesixseveneightnienzeroonetwothreefourfivesixseveneightnien")
	}
}
