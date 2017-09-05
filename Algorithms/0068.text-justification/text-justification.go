package Problem0068

func fullJustify(words []string, maxWidth int) []string {
	res := []string{}
	temp := []string{}
	width := 0

	for len(words) > 0 {
		words, temp, width = split(words, maxWidth)
		res = append(res, combine(temp, width, maxWidth))
	}

	return res
}

// 返回待组合的单词，需要组合的单词，和这些单词的长度之和
func split(words []string, maxWidth int) ([]string, []string, int) {
	res := make([]string, 1)
	res[0] = words[0]
	width := len(words[0])
	i := 1
	for ; i < len(words); i++ {
		if width+len(words[i]) > maxWidth {
			break
		}
		res = append(res, words[i])
		width += len(words[i])
	}
	return words[i:], res, width
}

func combine(words []string, width, maxWidth int) string {
	wordCount := len(words)
	spaces := makeSpaces(wordCount-1, maxWidth-width)

	res := ""
	for i, v := range spaces {
		res += words[i] + v
	}
	res += words[wordCount-1]

	return res
}

func makeSpaces(Len, count int) []string {
	res := make([]string, Len)
	for i := 0; i < count; i++ {
		res[i%Len] += " "
	}

	return res
}
