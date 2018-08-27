package problem0893

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []string
	ans int
}{

	{
		[]string{"demp", "cfhp", "dzvf", "ggxe", "hkte", "clug", "nhgk", "hvwj", "zozr", "datm", "hisr", "gfry", "jknr", "laju", "emsf", "duwe", "ilta", "gjrd", "woaq", "zhdm", "ujmz", "jalu", "tkhe", "gexg", "hrsi", "tail", "ilta", "xegg", "srhi", "clug", "cgul", "gexg", "tehk", "ulcg", "xgge", "cgul", "hrsi", "aowq", "jalu", "laju", "vzdf", "gexg", "hpcf", "zhdm", "hfcp", "zhdm", "ulcg", "jalu", "ggxe", "gexg", "nkgh", "hrsi", "vfdz", "nkgh", "cgul", "hpcf", "hetk", "zrzo", "xegg", "nkgh", "srhi", "smef", "ulcg", "hrsi", "ggxe", "ggxe", "efsm", "ggxe", "jalu", "srhi", "dmzh", "laju", "zmdh", "sfem", "tehk", "srhi", "wqao", "gknh", "jalu", "iatl", "gexg", "ugcl", "nkgh", "hrsi", "srhi", "hkte", "gdrj", "zozr", "hisr", "sihr", "smef", "zmdh", "clug", "iatl", "cgul", "woaq", "jrnk", "sihr", "xegg", "luja"},
		21,
	},

	{
		[]string{"a", "b", "c", "a", "c", "c"},
		3,
	},

	{
		[]string{"aa", "bb", "ab", "ba"},
		4,
	},

	{
		[]string{"abc", "acb", "bac", "bca", "cab", "cba"},
		3,
	},

	{
		[]string{"abcd", "cdab", "adcb", "cbad"},
		1,
	},

	// 可以有多个 testcase
}

func Test_numSpecialEquivGroups(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, numSpecialEquivGroups(tc.A), "输入:%v", tc)
	}
}

func Benchmark_numSpecialEquivGroups(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numSpecialEquivGroups(tc.A)
		}
	}
}
