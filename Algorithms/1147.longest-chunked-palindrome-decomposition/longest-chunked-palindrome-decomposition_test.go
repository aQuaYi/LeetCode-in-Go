package problem1147

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	text string
	ans  int
}{

	{
		"ykpwikcghnoxoumuggqqybujrbkmnjlccsvjpoglirbrpgwkwxlmnfcpveijqluopugmksmmrfblaciqqtyidjxdrrefhoszhywhzjlvxmvfcmtszcclkhzkarheolcwikciixvarpffcdserxpzfpmrlxgmahxonomqzjfhjznvzbvsmwxfalcsdbaknspktjvydykivtfwvzicnausanqabozwcxpmimmldzpeiujilsixyrvxiisaeataeeydfodjqbpbserygikmcczluabsujnkfgemcdszftwkyteeagthkspkvkehlvdxkjnwuwmoyhdyksybqoqwdbrrabhkxuolvxrnyopxsqucilfakiiuwvwnhxclxnwvgxqsevcgdwuiaqtqbbkwporlsgbnotnhbcumfzzzvldzlqyiyhdfgsdqyifzyqecyesuygxynosctshoohumujzmrwfzaxjcjtubzugtiwekrtlluudmlqooqldmtlluuiwekrubzugtxjcjtzafjzmrwoohumutshynoscuygxesyqecydqyifzfgsyiyhdqdzlzzvlcumfzbhnotngblsbkwporiaqtqbvcgdwuqsenwvgxlxvwnhxciuwfakiluciqopxsvxrnykxuolrabhdbrqoqwybdyksuwmoyhdxkjnwehlvkpkvkseagthtwkytemcdszfnkfgeuabsujzlikmccgbseryodjqbpydfaeesaeatvxiisixyreiujilpldzmimmcxpqabozwnausanvzicfwdykivtpktjvyknsaalcsdbsmwxfnvzbvjfhjzqzommahxonxgmrlpzfpdserxffcxvarpikciwiheolchzkarzcclkfcmtsjlvxmvhywhzzsrefhojxdrqtyidaciqfblmrgmksmpuluoveijqpnfcmlpgwkwxbroglirvjpjlccskmnybujrbmuggqqnoxouikcghykpw",
		288,
	},

	{
		"ghiabcdefhelloadamhelloabcdefghi",
		7,
	},

	{
		"merchant",
		1,
	},

	{
		"antaprezatepzapreanta",
		11,
	},

	{
		"aaa",
		3,
	},

	// 可以有多个 testcase
}

func Test_longestDecomposition(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, longestDecomposition(tc.text), "输入:%v", tc)
	}
}

func Benchmark_longestDecomposition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			longestDecomposition(tc.text)
		}
	}
}
