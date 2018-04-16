package problem0804

import (
	"bytes"
	"fmt"
)

var table = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations(words []string) int {
	res := make(map[string]bool, len(words))
	for _, w := range words {
		var b bytes.Buffer
		for i := 0; i < len(w); i++ {
			fmt.Fprint(&b, table[w[i]-'a'])
		}
		res[b.String()] = true
	}
	return len(res)
}

// 对于 go >= 1.10 来说，以下才是最快的方法

// func uniqueMorseRepresentations(words []string) int {
// 	res := make(map[string]bool, len(words))
// 	for _, w := range words {
// 		var b strings.Builder
// 		for i := 0; i < len(w); i++ {
// 			b.WriteString(table[w[i]-'a'])
// 		}
// 		res[b.String()] = true
// 	}
// 	return len(res)
// }
