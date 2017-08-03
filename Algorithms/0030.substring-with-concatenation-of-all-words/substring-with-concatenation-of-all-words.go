package Problem0030

func findSubstring(s string, words []string) []int {
	res := []int{}
	if len(s) == 0 || len(words) == 0 || len(s) < len(words)*len(words[0]) {
		return res
	}

	initialise := func(wordMap map[string]int, words []string) {
		for _, word := range words {
			wordMap[word] = 0
		}
		for _, word := range words {
			wordMap[word]++
		}
	}

	wordMap := make(map[string]int, len(words))
	lens, numWords, lenw := len(s), len(words), len(words[0])

	for trace := 0; trace < lenw; trace++ {
		initialise(wordMap, words)
		begin, count := trace, 0

		for begin+numWords*lenw <= lens {
			w := s[begin+count*lenw : begin+(count+1)*lenw]

			num, ok := wordMap[w]
			switch {
			case !ok:
				// 出现了words以外的单词
				// 从 w 后面一个字符，重新开始
				begin += lenw * (count + 1)
				if count != 0 {
					// wordMap 已经被修改，需要再次初始化
					initialise(wordMap, words)
					count = 0
				}
			case num == 0:
				// w 的出现次数超标了

				// 修改统计记录，从 begin 后的一个 word 开始记录
				// 增加 begin 处 word 可出现次数一次，
				wordMap[s[begin:begin+lenw]]++
				// 统计记录，增加一次
				count--
				// begin 后移一个 word 的长度
				begin += lenw

				// 这个case会连续执行，知道begin指向上一个 w 的后面一个字符
			default:
				// ok && num != 0

				// 修改统计记录
				wordMap[w]--
				count++

				if count == numWords {
					res = append(res, begin)
					wordMap[s[begin:begin+lenw]]++
					count--
					begin += lenw
				}
			}
		}
	}
	return res
}
