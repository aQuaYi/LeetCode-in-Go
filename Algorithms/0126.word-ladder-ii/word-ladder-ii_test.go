package Problem0126

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Problem0126(t *testing.T) {
	ast := assert.New(t)

	// tcs is testcase slice
	tcs := []struct {
		beginWord string
		endWord   string
		wordList  []string
		ans       [][]string
	}{

		{
			"qa",
			"sq",
			[]string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb", "sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb", "ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr", "sq", "ye"},
			[][]string{
				[]string{"qa", "ba", "be", "se", "sq"},
				[]string{"qa", "ba", "bi", "si", "sq"},
				[]string{"qa", "ba", "br", "sr", "sq"},
				[]string{"qa", "ca", "cm", "sm", "sq"},
				[]string{"qa", "ca", "co", "so", "sq"},
				[]string{"qa", "la", "ln", "sn", "sq"},
				[]string{"qa", "la", "lt", "st", "sq"},
				[]string{"qa", "ma", "mb", "sb", "sq"},
				[]string{"qa", "pa", "ph", "sh", "sq"},
				[]string{"qa", "ta", "tc", "sc", "sq"},
				[]string{"qa", "fa", "fe", "se", "sq"},
				[]string{"qa", "ga", "ge", "se", "sq"},
				[]string{"qa", "ha", "he", "se", "sq"},
				[]string{"qa", "la", "le", "se", "sq"},
				[]string{"qa", "ma", "me", "se", "sq"},
				[]string{"qa", "na", "ne", "se", "sq"},
				[]string{"qa", "ra", "re", "se", "sq"},
				[]string{"qa", "ya", "ye", "se", "sq"},
				[]string{"qa", "ca", "ci", "si", "sq"},
				[]string{"qa", "ha", "hi", "si", "sq"},
				[]string{"qa", "la", "li", "si", "sq"},
				[]string{"qa", "ma", "mi", "si", "sq"},
				[]string{"qa", "na", "ni", "si", "sq"},
				[]string{"qa", "pa", "pi", "si", "sq"},
				[]string{"qa", "ta", "ti", "si", "sq"},
				[]string{"qa", "ca", "cr", "sr", "sq"},
				[]string{"qa", "fa", "fr", "sr", "sq"},
				[]string{"qa", "la", "lr", "sr", "sq"},
				[]string{"qa", "ma", "mr", "sr", "sq"},
				[]string{"qa", "fa", "fm", "sm", "sq"},
				[]string{"qa", "pa", "pm", "sm", "sq"},
				[]string{"qa", "ta", "tm", "sm", "sq"},
				[]string{"qa", "ga", "go", "so", "sq"},
				[]string{"qa", "ha", "ho", "so", "sq"},
				[]string{"qa", "la", "lo", "so", "sq"},
				[]string{"qa", "ma", "mo", "so", "sq"},
				[]string{"qa", "na", "no", "so", "sq"},
				[]string{"qa", "pa", "po", "so", "sq"},
				[]string{"qa", "ta", "to", "so", "sq"},
				[]string{"qa", "ya", "yo", "so", "sq"},
				[]string{"qa", "ma", "mn", "sn", "sq"},
				[]string{"qa", "ra", "rn", "sn", "sq"},
				[]string{"qa", "ma", "mt", "st", "sq"},
				[]string{"qa", "pa", "pt", "st", "sq"},
				[]string{"qa", "na", "nb", "sb", "sq"},
				[]string{"qa", "pa", "pb", "sb", "sq"},
				[]string{"qa", "ra", "rb", "sb", "sq"},
				[]string{"qa", "ta", "tb", "sb", "sq"},
				[]string{"qa", "ya", "yb", "sb", "sq"},
				[]string{"qa", "ra", "rh", "sh", "sq"},
				[]string{"qa", "ta", "th", "sh", "sq"},
			},
		},

		{
			"hit",
			"cog",
			[]string{"hot", "dot", "dog", "lot", "log"},
			[][]string{},
		},

		{
			"hot",
			"dog",
			[]string{"hot", "dog", "dot"},
			[][]string{
				[]string{"hot", "dot", "dog"},
			},
		},

		{
			"hit",
			"cog",
			[]string{"hot", "dot", "dog", "lot", "log", "cog"},
			[][]string{
				[]string{"hit", "hot", "dot", "dog", "cog"},
				[]string{"hit", "hot", "lot", "log", "cog"},
			},
		},

		// 可以多个 testcase
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, findLadders(tc.beginWord, tc.endWord, tc.wordList), "输入:%v", tc)
	}
}
