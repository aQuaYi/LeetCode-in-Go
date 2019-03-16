package problem0953

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	words []string
	order string
	ans   bool
}{

	{
		[]string{"skuxofhyqffmvem", "pivlrvyixuwkteq", "alsgqypugsbbobl", "aaeognfoaxhcfii", "mpyonsmytxdnysa", "midsibtcmyeshoy", "ozqieboykydhjlz", "bolwcxirutfkovw", "kmbgtngnzhldnql", "xrxbriikmpxrprn", "xasnbmzcyncynqf", "srvryjpvrkjfwum", "kxcytrfssxrrksl", "foiaripeubtyyox", "yyhnjykzexgksob", "tykvetxwxbnaghl", "pkfmfhnboutbzry", "mlxxxcpeapjhwvx", "jmtcevkbcurjqhp", "hpowpmxkjiufoyz", "trbrldyxjbthcwi", "elkqzwcdljzzilp", "xuututciwchdtgp", "dznknlxbhjqlnfp", "usvebdafpehvhnn", "vgztrhsazakvkoy", "ppeowsoiwauodrd", "rwfnkmwpowiilmn", "gnptotnrcujhobx", "rfesptvhuiwtzuv", "nncoygoxusmskdj", "zrdmgclmxuygrfa", "ccuprnhgisbhnlo", "iwwnyuooaxciddu", "kyyzpnkywltfyqf", "fwvfdgyhgwnedww", "inleqhipjuvaipy", "ggwajwbrmbsvzto", "smfnzvhxnzlttqq", "hgolraueyiveyff", "bcitwifwgcvnfiz", "tporichlcybalot", "axjnikarmsedfkf", "takwnduamciroyg", "jtoikpsttsdiusp", "tlwzeycxdkigmia", "gvktgerybyghsoi", "gexklhkdamzxsar", "eaescmsyvbfhjki", "izkzehqhdpfkcfd", "riygtdcgibpigjr", "hxnngtftblikoan", "hqeyivtoimdmcgw", "srgbehdjwgvkwfd", "syrqpdzhbwrohvw", "ulghcdyoaaimdsk", "yotgykrwulptddt", "vmujiultrruvicw", "bvgazhtfaxopncl", "tlygnypwvdpqruq", "gwghjruwprmaywb", "tmjlyehemrjllsf", "jbtollpqxhcxipc", "urryhtrjdcyyxgk", "fovuiicvzbvopje", "iocnnebdisktpto", "dizpowvljuxcuyp", "lmbryijbblvbfew", "hsbdhkhaqjsyezo", "xuzrakcohgkzvja", "rqytcuxmbexynso", "muvxuufnbdxpgqp", "kazzhqfnoerazdp", "ydcbxbysbzqavgq", "uvswbtjojzodhxx", "ogficmoxohwmacf", "rsxarauxrlsugzl", "ivxddltcdfqnsku", "flxhjxcbldrhmtg", "gwcgyybialciiaz", "euizzqwnnefihcz", "ttlzrgnwhgzkirj", "geomyyrdrhkimzv", "wfoxpjisclyoygt", "iunuvuqdkeqqacf", "vfftvmtoaanoafp", "ogzwqioazjyedjq", "iltzlygzsreqlkw", "paighhiwamnafai", "aslchefeszbmokl", "opvtsdbqxgppvmj", "kfjaofmhvgfjxja", "xiwykodfpgizgky", "qsznrasrvabazev", "qpevsngotolecsi"},
		"fhgbavxiyjpueqkodmtzncslwr",
		false,
	},

	{
		[]string{"hello", "leetcode"},
		"hlabcdefgijkmnopqrstuvwxyz",
		true,
	},

	{
		[]string{"word", "world", "row"},
		"worldabcefghijkmnpqstuvxyz",
		false,
	},

	{
		[]string{"apple", "app"},
		"abcdefghijklmnopqrstuvwxyz",
		false,
	},

	// 可以有多个 testcase
}

func Test_isAlienSorted(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, isAlienSorted(tc.words, tc.order), "输入:%v", tc)
	}
}

func Benchmark_isAlienSorted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isAlienSorted(tc.words, tc.order)
		}
	}
}
