package problem0068

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	res := []string{}
	temp := []string{}
	width := 0
	isLast := false

	for !isLast {
		words, temp, width, isLast = split(words, maxWidth)
		res = append(res, combine(temp, width, maxWidth, isLast))
	}

	return res
}

// 返回不组合的单词，需要组合的单词，和这些单词的长度之和, 是否是结尾
func split(words []string, maxWidth int) ([]string, []string, int, bool) {
	temp := make([]string, 1)
	temp[0] = words[0]
	width := len(words[0])

	i := 1
	for ; i < len(words); i++ {
		if width+len(temp)+len(words[i]) > maxWidth {
			break
		}
		temp = append(temp, words[i])
		width += len(words[i])
	}

	return words[i:], temp, width, i == len(words)
}

func combine(words []string, width, maxWidth int, isLast bool) string {
	wordCount := len(words)
	if wordCount == 1 || isLast {
		return combineSpecial(words, maxWidth)
	}
	spaceCount := wordCount - 1
	spaces := makeSpaces(spaceCount, maxWidth-width)

	res := ""
	for i, v := range spaces {
		res += words[i] + v
	}
	if wordCount > 1 {
		res += words[wordCount-1]
	}

	return res
}

func makeSpaces(Len, count int) []string {
	res := make([]string, Len)
	for i := 0; i < count; i++ {
		res[i%Len] += " "
	}

	return res
}

func combineSpecial(words []string, maxWidth int) string {
	res := strings.Join(words, " ")

	for len(res) < maxWidth {
		res += " "
	}

	return res
}
