package problem1081

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	text string
	ans  string
}{

	{
		"fduxqmmeghmsrohvpclyjmhwwxipnvqxhfmgtqtcutsbjhbmuollkwocxxdmoswsnksdflwbxetsvbuvlwtxonpyhyrodjmdbvoopfxinkojyugqegbprgqxjpghojhymsoqpylrmelsonpqrtjmjgbgmmpqklfiiaacgurrbqtbqyylfqiefbrfyotiptqujknegwjiyqybldgomccdbiikfsfnqwcilblilfwcxyytnvdrppmcslildixlungoetlqvpvpwgmhqvwwjmllomtipfavbhbahclcfdyvgyqhpxebmtovgxqtjwdiwkqtvnlumwjgvubghkcjsvkrydpasdknkdclutjcqbretopqobivwfdkqkvmwkkufwnrngfgixlinerxcnrmsbiybcxmmndhhdrwykwmgckxqhlhnabppswwkxbjpvpeplcyyhcvhjulxvvgabddcurghjurledjdatsytdkqlfyrpnasrqiyecvjtkoiuigawvqfemwwnpkhapxvaqrlnncxdepunrnimqwcinnbnifvsjkwhufoawtbeghauvxiggajubybemfyeropjwvuhjrtiggsoaddpbgfcftppwnnlgnhbrdbhycslqlfkwdiswxntapahkpsyufkthkgmbvtmbnutyhrpjhotpndnldiugmmgxmtsdxqjojaqbedotgxlgaqyempwjlvtgifybqmxvcfbuonwivfhpqrmxfakrhjrsxovxgfwcteuadldgyghvcrjaaomwisvouyqqdmdbsbuvepcaxtkuqtsvqjbmejvptmbmbxasbxadvauepxicyjnydhmsvlohnnqoevhewwmxuqingvgbniqouikflimmxpuygutamkthmkydlwtvigyeutikpvnoqisehcgaylunxgwdvanwhsbjkukxahuviutenigfmkblniwdtqxxnswjoxhqqooaltrubgbqvqqdmdmrixidgukqx",
		"abcdefghijklmnopqrstuvwxy",
	},

	{
		"cdadabcc",
		"adbc",
	},

	{
		"abcd",
		"abcd",
	},

	{
		"ecbacba",
		"eacb",
	},

	{
		"leetcode",
		"letcod",
	},

	// 可以有多个 testcase
}

func Test_smallestSubsequence(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, smallestSubsequence(tc.text), "输入:%v", tc)
	}
}

func Benchmark_smallestSubsequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			smallestSubsequence(tc.text)
		}
	}
}
