package problem0316

import (
	"strings"
)

func removeDuplicateLetters(s string) string {
	lens := len(s)
	if lens == 0 {
		return ""
	}
	// 统计每个字母出现的次数
	count := [26]int{}
	var ch rune
	for _, ch = range s {
		count[ch-'a']++
	}

	// index of min character
	imc := 0
	for i := range s {
		if s[imc] > s[i] {
			imc = i
		}
		count[s[i]-'a']--
		if count[s[i]-'a'] == 0 {
			break
		}
	}

	// 贪心算法：更小的 s[i] 值得拥有更小的 i
	// 1. s[:ims] 中出现的字母，在 s[ims+1:] 中都会再出现
	// 2. s[ims] 是 s[:ims+1] 中最小的字母
	// 综合以上两条
	// 1. 删除 s[:ims]
	// 2. 把 s[ims] 作为答案的首字母
	// 3. 删除 s[ims+1:] 中的 s[ims] 后，递归求解
	return string(s[imc]) + removeDuplicateLetters(strings.Replace(s[imc+1:], string(s[imc]), "", -1))
}
