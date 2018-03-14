package problem0767

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
		"abbabbaaab",
		"ababababab",
	},

	{
		"snnnnbpngobwznvnnnlnwhvnnnnfjnnlnnnnnnbnknnqkndzefncknnnnnaiqrntnndnnnjninnnunnunqhndnnqnnsjqnnpiqshntnnncnvnnnncnnqenlnninyndnnnljongnnjwnnnngllnnngkbnllnnnnontlbpngjnnenqnsnnnnnjeqqghnfpngepnodnnnnnnvnsrnughbnipvnhqmnzonoonnnjotnnonoennnpnfnnkdnnbmnmnpnqninnxronnnnvnlanlnnnebnnnlnvnfknsnbincnttnmnguqenhnnxunnnntnnnnhnqnzehvunfnvnndvnjnnnbnnpxnqipwnmnonnndlnsnonnninnxnnnjnnnnnesennmyiednnnnnnnnnhimtnnnonjlicnwnwvnntaxmnrntnnnnsnbnanninnecbcfjxncnnkvnnqgnunensanpnngjnzxjnopnnyvnnxskniyytnsnnnnx",
		"nqnqnqnqnqnqnqnqnqnqnqnqnqnqnqnqnonononononononononononononononenenenenenenenenenenenenenenininininininininininininininjnjnjnjnjnjnjnjnjnjnjnjnjnjnlnlnlnlnlnlnlnlnlnlnlnlnlnlnvnvnvnvnvnvnvnvnvnvnvnvnvnbnbnbnbnbnbnbnbnbnbnbnbnpnpnpnpnpnpnpnpnpnpnpnpnsnsnsnsnsnsnsnsnsnsnsnsngngngngngngngngngngngntntntntntntntntntntntnhnhnhnhnhnhnhnhnhnhndndndndndndndndndnxnxnxnxnxnxnxnxnxncncncncncncncncnknknknknknknknknmnmnmnmnmnmnmnmnfnfnfnfnfnfnfnunununununununwnwnwnwnwnwnanananananynynynynynznznznznznrnrnrnrn",
	},

	{
		"aaab",
		"",
	},

	{
		"aab",
		"aba",
	},

	{
		"baaba",
		"ababa",
	},

	{
		"vvvlo",
		"vlvov",
	},

	// 可以有多个 testcase
}

func Test_reorganizeString(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, reorganizeString(tc.S), "输入:%v", tc)
	}
}

func Benchmark_reorganizeString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			reorganizeString(tc.S)
		}
	}
}
