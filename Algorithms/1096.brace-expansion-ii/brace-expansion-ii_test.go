package problem1096

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	expression string
	ans        []string
}{

	{
		"{a,b}{a,b}{a,b}{a,b}{a,b}{a,b}{a,b}{a,b}{a,b}{a,b}",
		[]string{"aaaaaaaaaa", "aaaaaaaaab", "aaaaaaaaba", "aaaaaaaabb", "aaaaaaabaa", "aaaaaaabab", "aaaaaaabba", "aaaaaaabbb", "aaaaaabaaa", "aaaaaabaab", "aaaaaababa", "aaaaaababb", "aaaaaabbaa", "aaaaaabbab", "aaaaaabbba", "aaaaaabbbb", "aaaaabaaaa", "aaaaabaaab", "aaaaabaaba", "aaaaabaabb", "aaaaababaa", "aaaaababab", "aaaaababba", "aaaaababbb", "aaaaabbaaa", "aaaaabbaab", "aaaaabbaba", "aaaaabbabb", "aaaaabbbaa", "aaaaabbbab", "aaaaabbbba", "aaaaabbbbb", "aaaabaaaaa", "aaaabaaaab", "aaaabaaaba", "aaaabaaabb", "aaaabaabaa", "aaaabaabab", "aaaabaabba", "aaaabaabbb", "aaaababaaa", "aaaababaab", "aaaabababa", "aaaabababb", "aaaababbaa", "aaaababbab", "aaaababbba", "aaaababbbb", "aaaabbaaaa", "aaaabbaaab", "aaaabbaaba", "aaaabbaabb", "aaaabbabaa", "aaaabbabab", "aaaabbabba", "aaaabbabbb", "aaaabbbaaa", "aaaabbbaab", "aaaabbbaba", "aaaabbbabb", "aaaabbbbaa", "aaaabbbbab", "aaaabbbbba", "aaaabbbbbb", "aaabaaaaaa", "aaabaaaaab", "aaabaaaaba", "aaabaaaabb", "aaabaaabaa", "aaabaaabab", "aaabaaabba", "aaabaaabbb", "aaabaabaaa", "aaabaabaab", "aaabaababa", "aaabaababb", "aaabaabbaa", "aaabaabbab", "aaabaabbba", "aaabaabbbb", "aaababaaaa", "aaababaaab", "aaababaaba", "aaababaabb", "aaabababaa", "aaabababab", "aaabababba", "aaabababbb", "aaababbaaa", "aaababbaab", "aaababbaba", "aaababbabb", "aaababbbaa", "aaababbbab", "aaababbbba", "aaababbbbb", "aaabbaaaaa", "aaabbaaaab", "aaabbaaaba", "aaabbaaabb", "aaabbaabaa", "aaabbaabab", "aaabbaabba", "aaabbaabbb", "aaabbabaaa", "aaabbabaab", "aaabbababa", "aaabbababb", "aaabbabbaa", "aaabbabbab", "aaabbabbba", "aaabbabbbb", "aaabbbaaaa", "aaabbbaaab", "aaabbbaaba", "aaabbbaabb", "aaabbbabaa", "aaabbbabab", "aaabbbabba", "aaabbbabbb", "aaabbbbaaa", "aaabbbbaab", "aaabbbbaba", "aaabbbbabb", "aaabbbbbaa", "aaabbbbbab", "aaabbbbbba", "aaabbbbbbb", "aabaaaaaaa", "aabaaaaaab", "aabaaaaaba", "aabaaaaabb", "aabaaaabaa", "aabaaaabab", "aabaaaabba", "aabaaaabbb", "aabaaabaaa", "aabaaabaab", "aabaaababa", "aabaaababb", "aabaaabbaa", "aabaaabbab", "aabaaabbba", "aabaaabbbb", "aabaabaaaa", "aabaabaaab", "aabaabaaba", "aabaabaabb", "aabaababaa", "aabaababab", "aabaababba", "aabaababbb", "aabaabbaaa", "aabaabbaab", "aabaabbaba", "aabaabbabb", "aabaabbbaa", "aabaabbbab", "aabaabbbba", "aabaabbbbb", "aababaaaaa", "aababaaaab", "aababaaaba", "aababaaabb", "aababaabaa", "aababaabab", "aababaabba", "aababaabbb", "aabababaaa", "aabababaab", "aababababa", "aababababb", "aabababbaa", "aabababbab", "aabababbba", "aabababbbb", "aababbaaaa", "aababbaaab", "aababbaaba", "aababbaabb", "aababbabaa", "aababbabab", "aababbabba", "aababbabbb", "aababbbaaa", "aababbbaab", "aababbbaba", "aababbbabb", "aababbbbaa", "aababbbbab", "aababbbbba", "aababbbbbb", "aabbaaaaaa", "aabbaaaaab", "aabbaaaaba", "aabbaaaabb", "aabbaaabaa", "aabbaaabab", "aabbaaabba", "aabbaaabbb", "aabbaabaaa", "aabbaabaab", "aabbaababa", "aabbaababb", "aabbaabbaa", "aabbaabbab", "aabbaabbba", "aabbaabbbb", "aabbabaaaa", "aabbabaaab", "aabbabaaba", "aabbabaabb", "aabbababaa", "aabbababab", "aabbababba", "aabbababbb", "aabbabbaaa", "aabbabbaab", "aabbabbaba", "aabbabbabb", "aabbabbbaa", "aabbabbbab", "aabbabbbba", "aabbabbbbb", "aabbbaaaaa", "aabbbaaaab", "aabbbaaaba", "aabbbaaabb", "aabbbaabaa", "aabbbaabab", "aabbbaabba", "aabbbaabbb", "aabbbabaaa", "aabbbabaab", "aabbbababa", "aabbbababb", "aabbbabbaa", "aabbbabbab", "aabbbabbba", "aabbbabbbb", "aabbbbaaaa", "aabbbbaaab", "aabbbbaaba", "aabbbbaabb", "aabbbbabaa", "aabbbbabab", "aabbbbabba", "aabbbbabbb", "aabbbbbaaa", "aabbbbbaab", "aabbbbbaba", "aabbbbbabb", "aabbbbbbaa", "aabbbbbbab", "aabbbbbbba", "aabbbbbbbb", "abaaaaaaaa", "abaaaaaaab", "abaaaaaaba", "abaaaaaabb", "abaaaaabaa", "abaaaaabab", "abaaaaabba", "abaaaaabbb", "abaaaabaaa", "abaaaabaab", "abaaaababa", "abaaaababb", "abaaaabbaa", "abaaaabbab", "abaaaabbba", "abaaaabbbb", "abaaabaaaa", "abaaabaaab", "abaaabaaba", "abaaabaabb", "abaaababaa", "abaaababab", "abaaababba", "abaaababbb", "abaaabbaaa", "abaaabbaab", "abaaabbaba", "abaaabbabb", "abaaabbbaa", "abaaabbbab", "abaaabbbba", "abaaabbbbb", "abaabaaaaa", "abaabaaaab", "abaabaaaba", "abaabaaabb", "abaabaabaa", "abaabaabab", "abaabaabba", "abaabaabbb", "abaababaaa", "abaababaab", "abaabababa", "abaabababb", "abaababbaa", "abaababbab", "abaababbba", "abaababbbb", "abaabbaaaa", "abaabbaaab", "abaabbaaba", "abaabbaabb", "abaabbabaa", "abaabbabab", "abaabbabba", "abaabbabbb", "abaabbbaaa", "abaabbbaab", "abaabbbaba", "abaabbbabb", "abaabbbbaa", "abaabbbbab", "abaabbbbba", "abaabbbbbb", "ababaaaaaa", "ababaaaaab", "ababaaaaba", "ababaaaabb", "ababaaabaa", "ababaaabab", "ababaaabba", "ababaaabbb", "ababaabaaa", "ababaabaab", "ababaababa", "ababaababb", "ababaabbaa", "ababaabbab", "ababaabbba", "ababaabbbb", "abababaaaa", "abababaaab", "abababaaba", "abababaabb", "ababababaa", "ababababab", "ababababba", "ababababbb", "abababbaaa", "abababbaab", "abababbaba", "abababbabb", "abababbbaa", "abababbbab", "abababbbba", "abababbbbb", "ababbaaaaa", "ababbaaaab", "ababbaaaba", "ababbaaabb", "ababbaabaa", "ababbaabab", "ababbaabba", "ababbaabbb", "ababbabaaa", "ababbabaab", "ababbababa", "ababbababb", "ababbabbaa", "ababbabbab", "ababbabbba", "ababbabbbb", "ababbbaaaa", "ababbbaaab", "ababbbaaba", "ababbbaabb", "ababbbabaa", "ababbbabab", "ababbbabba", "ababbbabbb", "ababbbbaaa", "ababbbbaab", "ababbbbaba", "ababbbbabb", "ababbbbbaa", "ababbbbbab", "ababbbbbba", "ababbbbbbb", "abbaaaaaaa", "abbaaaaaab", "abbaaaaaba", "abbaaaaabb", "abbaaaabaa", "abbaaaabab", "abbaaaabba", "abbaaaabbb", "abbaaabaaa", "abbaaabaab", "abbaaababa", "abbaaababb", "abbaaabbaa", "abbaaabbab", "abbaaabbba", "abbaaabbbb", "abbaabaaaa", "abbaabaaab", "abbaabaaba", "abbaabaabb", "abbaababaa", "abbaababab", "abbaababba", "abbaababbb", "abbaabbaaa", "abbaabbaab", "abbaabbaba", "abbaabbabb", "abbaabbbaa", "abbaabbbab", "abbaabbbba", "abbaabbbbb", "abbabaaaaa", "abbabaaaab", "abbabaaaba", "abbabaaabb", "abbabaabaa", "abbabaabab", "abbabaabba", "abbabaabbb", "abbababaaa", "abbababaab", "abbabababa", "abbabababb", "abbababbaa", "abbababbab", "abbababbba", "abbababbbb", "abbabbaaaa", "abbabbaaab", "abbabbaaba", "abbabbaabb", "abbabbabaa", "abbabbabab", "abbabbabba", "abbabbabbb", "abbabbbaaa", "abbabbbaab", "abbabbbaba", "abbabbbabb", "abbabbbbaa", "abbabbbbab", "abbabbbbba", "abbabbbbbb", "abbbaaaaaa", "abbbaaaaab", "abbbaaaaba", "abbbaaaabb", "abbbaaabaa", "abbbaaabab", "abbbaaabba", "abbbaaabbb", "abbbaabaaa", "abbbaabaab", "abbbaababa", "abbbaababb", "abbbaabbaa", "abbbaabbab", "abbbaabbba", "abbbaabbbb", "abbbabaaaa", "abbbabaaab", "abbbabaaba", "abbbabaabb", "abbbababaa", "abbbababab", "abbbababba", "abbbababbb", "abbbabbaaa", "abbbabbaab", "abbbabbaba", "abbbabbabb", "abbbabbbaa", "abbbabbbab", "abbbabbbba", "abbbabbbbb", "abbbbaaaaa", "abbbbaaaab", "abbbbaaaba", "abbbbaaabb", "abbbbaabaa", "abbbbaabab", "abbbbaabba", "abbbbaabbb", "abbbbabaaa", "abbbbabaab", "abbbbababa", "abbbbababb", "abbbbabbaa", "abbbbabbab", "abbbbabbba", "abbbbabbbb", "abbbbbaaaa", "abbbbbaaab", "abbbbbaaba", "abbbbbaabb", "abbbbbabaa", "abbbbbabab", "abbbbbabba", "abbbbbabbb", "abbbbbbaaa", "abbbbbbaab", "abbbbbbaba", "abbbbbbabb", "abbbbbbbaa", "abbbbbbbab", "abbbbbbbba", "abbbbbbbbb", "baaaaaaaaa", "baaaaaaaab", "baaaaaaaba", "baaaaaaabb", "baaaaaabaa", "baaaaaabab", "baaaaaabba", "baaaaaabbb", "baaaaabaaa", "baaaaabaab", "baaaaababa", "baaaaababb", "baaaaabbaa", "baaaaabbab", "baaaaabbba", "baaaaabbbb", "baaaabaaaa", "baaaabaaab", "baaaabaaba", "baaaabaabb", "baaaababaa", "baaaababab", "baaaababba", "baaaababbb", "baaaabbaaa", "baaaabbaab", "baaaabbaba", "baaaabbabb", "baaaabbbaa", "baaaabbbab", "baaaabbbba", "baaaabbbbb", "baaabaaaaa", "baaabaaaab", "baaabaaaba", "baaabaaabb", "baaabaabaa", "baaabaabab", "baaabaabba", "baaabaabbb", "baaababaaa", "baaababaab", "baaabababa", "baaabababb", "baaababbaa", "baaababbab", "baaababbba", "baaababbbb", "baaabbaaaa", "baaabbaaab", "baaabbaaba", "baaabbaabb", "baaabbabaa", "baaabbabab", "baaabbabba", "baaabbabbb", "baaabbbaaa", "baaabbbaab", "baaabbbaba", "baaabbbabb", "baaabbbbaa", "baaabbbbab", "baaabbbbba", "baaabbbbbb", "baabaaaaaa", "baabaaaaab", "baabaaaaba", "baabaaaabb", "baabaaabaa", "baabaaabab", "baabaaabba", "baabaaabbb", "baabaabaaa", "baabaabaab", "baabaababa", "baabaababb", "baabaabbaa", "baabaabbab", "baabaabbba", "baabaabbbb", "baababaaaa", "baababaaab", "baababaaba", "baababaabb", "baabababaa", "baabababab", "baabababba", "baabababbb", "baababbaaa", "baababbaab", "baababbaba", "baababbabb", "baababbbaa", "baababbbab", "baababbbba", "baababbbbb", "baabbaaaaa", "baabbaaaab", "baabbaaaba", "baabbaaabb", "baabbaabaa", "baabbaabab", "baabbaabba", "baabbaabbb", "baabbabaaa", "baabbabaab", "baabbababa", "baabbababb", "baabbabbaa", "baabbabbab", "baabbabbba", "baabbabbbb", "baabbbaaaa", "baabbbaaab", "baabbbaaba", "baabbbaabb", "baabbbabaa", "baabbbabab", "baabbbabba", "baabbbabbb", "baabbbbaaa", "baabbbbaab", "baabbbbaba", "baabbbbabb", "baabbbbbaa", "baabbbbbab", "baabbbbbba", "baabbbbbbb", "babaaaaaaa", "babaaaaaab", "babaaaaaba", "babaaaaabb", "babaaaabaa", "babaaaabab", "babaaaabba", "babaaaabbb", "babaaabaaa", "babaaabaab", "babaaababa", "babaaababb", "babaaabbaa", "babaaabbab", "babaaabbba", "babaaabbbb", "babaabaaaa", "babaabaaab", "babaabaaba", "babaabaabb", "babaababaa", "babaababab", "babaababba", "babaababbb", "babaabbaaa", "babaabbaab", "babaabbaba", "babaabbabb", "babaabbbaa", "babaabbbab", "babaabbbba", "babaabbbbb", "bababaaaaa", "bababaaaab", "bababaaaba", "bababaaabb", "bababaabaa", "bababaabab", "bababaabba", "bababaabbb", "babababaaa", "babababaab", "bababababa", "bababababb", "babababbaa", "babababbab", "babababbba", "babababbbb", "bababbaaaa", "bababbaaab", "bababbaaba", "bababbaabb", "bababbabaa", "bababbabab", "bababbabba", "bababbabbb", "bababbbaaa", "bababbbaab", "bababbbaba", "bababbbabb", "bababbbbaa", "bababbbbab", "bababbbbba", "bababbbbbb", "babbaaaaaa", "babbaaaaab", "babbaaaaba", "babbaaaabb", "babbaaabaa", "babbaaabab", "babbaaabba", "babbaaabbb", "babbaabaaa", "babbaabaab", "babbaababa", "babbaababb", "babbaabbaa", "babbaabbab", "babbaabbba", "babbaabbbb", "babbabaaaa", "babbabaaab", "babbabaaba", "babbabaabb", "babbababaa", "babbababab", "babbababba", "babbababbb", "babbabbaaa", "babbabbaab", "babbabbaba", "babbabbabb", "babbabbbaa", "babbabbbab", "babbabbbba", "babbabbbbb", "babbbaaaaa", "babbbaaaab", "babbbaaaba", "babbbaaabb", "babbbaabaa", "babbbaabab", "babbbaabba", "babbbaabbb", "babbbabaaa", "babbbabaab", "babbbababa", "babbbababb", "babbbabbaa", "babbbabbab", "babbbabbba", "babbbabbbb", "babbbbaaaa", "babbbbaaab", "babbbbaaba", "babbbbaabb", "babbbbabaa", "babbbbabab", "babbbbabba", "babbbbabbb", "babbbbbaaa", "babbbbbaab", "babbbbbaba", "babbbbbabb", "babbbbbbaa", "babbbbbbab", "babbbbbbba", "babbbbbbbb", "bbaaaaaaaa", "bbaaaaaaab", "bbaaaaaaba", "bbaaaaaabb", "bbaaaaabaa", "bbaaaaabab", "bbaaaaabba", "bbaaaaabbb", "bbaaaabaaa", "bbaaaabaab", "bbaaaababa", "bbaaaababb", "bbaaaabbaa", "bbaaaabbab", "bbaaaabbba", "bbaaaabbbb", "bbaaabaaaa", "bbaaabaaab", "bbaaabaaba", "bbaaabaabb", "bbaaababaa", "bbaaababab", "bbaaababba", "bbaaababbb", "bbaaabbaaa", "bbaaabbaab", "bbaaabbaba", "bbaaabbabb", "bbaaabbbaa", "bbaaabbbab", "bbaaabbbba", "bbaaabbbbb", "bbaabaaaaa", "bbaabaaaab", "bbaabaaaba", "bbaabaaabb", "bbaabaabaa", "bbaabaabab", "bbaabaabba", "bbaabaabbb", "bbaababaaa", "bbaababaab", "bbaabababa", "bbaabababb", "bbaababbaa", "bbaababbab", "bbaababbba", "bbaababbbb", "bbaabbaaaa", "bbaabbaaab", "bbaabbaaba", "bbaabbaabb", "bbaabbabaa", "bbaabbabab", "bbaabbabba", "bbaabbabbb", "bbaabbbaaa", "bbaabbbaab", "bbaabbbaba", "bbaabbbabb", "bbaabbbbaa", "bbaabbbbab", "bbaabbbbba", "bbaabbbbbb", "bbabaaaaaa", "bbabaaaaab", "bbabaaaaba", "bbabaaaabb", "bbabaaabaa", "bbabaaabab", "bbabaaabba", "bbabaaabbb", "bbabaabaaa", "bbabaabaab", "bbabaababa", "bbabaababb", "bbabaabbaa", "bbabaabbab", "bbabaabbba", "bbabaabbbb", "bbababaaaa", "bbababaaab", "bbababaaba", "bbababaabb", "bbabababaa", "bbabababab", "bbabababba", "bbabababbb", "bbababbaaa", "bbababbaab", "bbababbaba", "bbababbabb", "bbababbbaa", "bbababbbab", "bbababbbba", "bbababbbbb", "bbabbaaaaa", "bbabbaaaab", "bbabbaaaba", "bbabbaaabb", "bbabbaabaa", "bbabbaabab", "bbabbaabba", "bbabbaabbb", "bbabbabaaa", "bbabbabaab", "bbabbababa", "bbabbababb", "bbabbabbaa", "bbabbabbab", "bbabbabbba", "bbabbabbbb", "bbabbbaaaa", "bbabbbaaab", "bbabbbaaba", "bbabbbaabb", "bbabbbabaa", "bbabbbabab", "bbabbbabba", "bbabbbabbb", "bbabbbbaaa", "bbabbbbaab", "bbabbbbaba", "bbabbbbabb", "bbabbbbbaa", "bbabbbbbab", "bbabbbbbba", "bbabbbbbbb", "bbbaaaaaaa", "bbbaaaaaab", "bbbaaaaaba", "bbbaaaaabb", "bbbaaaabaa", "bbbaaaabab", "bbbaaaabba", "bbbaaaabbb", "bbbaaabaaa", "bbbaaabaab", "bbbaaababa", "bbbaaababb", "bbbaaabbaa", "bbbaaabbab", "bbbaaabbba", "bbbaaabbbb", "bbbaabaaaa", "bbbaabaaab", "bbbaabaaba", "bbbaabaabb", "bbbaababaa", "bbbaababab", "bbbaababba", "bbbaababbb", "bbbaabbaaa", "bbbaabbaab", "bbbaabbaba", "bbbaabbabb", "bbbaabbbaa", "bbbaabbbab", "bbbaabbbba", "bbbaabbbbb", "bbbabaaaaa", "bbbabaaaab", "bbbabaaaba", "bbbabaaabb", "bbbabaabaa", "bbbabaabab", "bbbabaabba", "bbbabaabbb", "bbbababaaa", "bbbababaab", "bbbabababa", "bbbabababb", "bbbababbaa", "bbbababbab", "bbbababbba", "bbbababbbb", "bbbabbaaaa", "bbbabbaaab", "bbbabbaaba", "bbbabbaabb", "bbbabbabaa", "bbbabbabab", "bbbabbabba", "bbbabbabbb", "bbbabbbaaa", "bbbabbbaab", "bbbabbbaba", "bbbabbbabb", "bbbabbbbaa", "bbbabbbbab", "bbbabbbbba", "bbbabbbbbb", "bbbbaaaaaa", "bbbbaaaaab", "bbbbaaaaba", "bbbbaaaabb", "bbbbaaabaa", "bbbbaaabab", "bbbbaaabba", "bbbbaaabbb", "bbbbaabaaa", "bbbbaabaab", "bbbbaababa", "bbbbaababb", "bbbbaabbaa", "bbbbaabbab", "bbbbaabbba", "bbbbaabbbb", "bbbbabaaaa", "bbbbabaaab", "bbbbabaaba", "bbbbabaabb", "bbbbababaa", "bbbbababab", "bbbbababba", "bbbbababbb", "bbbbabbaaa", "bbbbabbaab", "bbbbabbaba", "bbbbabbabb", "bbbbabbbaa", "bbbbabbbab", "bbbbabbbba", "bbbbabbbbb", "bbbbbaaaaa", "bbbbbaaaab", "bbbbbaaaba", "bbbbbaaabb", "bbbbbaabaa", "bbbbbaabab", "bbbbbaabba", "bbbbbaabbb", "bbbbbabaaa", "bbbbbabaab", "bbbbbababa", "bbbbbababb", "bbbbbabbaa", "bbbbbabbab", "bbbbbabbba", "bbbbbabbbb", "bbbbbbaaaa", "bbbbbbaaab", "bbbbbbaaba", "bbbbbbaabb", "bbbbbbabaa", "bbbbbbabab", "bbbbbbabba", "bbbbbbabbb", "bbbbbbbaaa", "bbbbbbbaab", "bbbbbbbaba", "bbbbbbbabb", "bbbbbbbbaa", "bbbbbbbbab", "bbbbbbbbba", "bbbbbbbbbb"},
	},

	{
		"a{b,c}{d,e}f{g,h},z",
		[]string{"abdfg", "abdfh", "abefg", "abefh", "acdfg", "acdfh", "acefg", "acefh", "z"},
	},

	{
		"{{a,b},{b,c}}",
		[]string{"a", "b", "c"},
	},

	{
		"a",
		[]string{"a"},
	},

	{
		"w",
		[]string{"w"},
	},

	{
		"{a,b,c}",
		[]string{"a", "b", "c"},
	},

	{
		"{a,b}{c,d}",
		[]string{"ac", "ad", "bc", "bd"},
	},

	{
		"a{b,c}{d,e}f{g,h}",
		[]string{"abdfg", "abdfh", "abefg", "abefh", "acdfg", "acdfh", "acefg", "acefh"},
	},

	{
		"ab{{{c}}}",
		[]string{"abc"},
	},

	{
		"{a,b}{c,{d,e}}",
		[]string{"ac", "ad", "ae", "bc", "bd", "be"},
	},

	{
		"{{a,z},a{b,c},{ab,z}}",
		[]string{"a", "ab", "ac", "z"},
	},

	// 可以有多个 testcase
}

func Test_braceExpansionII(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, braceExpansionII(tc.expression), "输入:%v", tc)
	}
}

func Benchmark_braceExpansionII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			braceExpansionII(tc.expression)
		}
	}
}
