package Problem0030

func findSubstring(s string, words []string) []int {
	res := []int{}

	lens, numWords, lenw := len(s), len(words), len(words[0])
	if lens == 0 || numWords == 0 || lens < numWords*lenw {
		return res
	}

	// remainNum 记录 words 中每个单词还能出现的次数
	remainNum := make(map[string]int, numWords)
	// count 记录符号要求的单词的连续出现次数
	count := 0
	initRecord := func() {
		for _, word := range words {
			remainNum[word] = 0
		}
		for _, word := range words {
			remainNum[word]++
		}
		count = 0
	} // 把这个匿名函数放在这里，可以避免繁琐的函数参数

	// 由于 index 增加的跨度是 lenw
	// index 需要分别从{0,1,...,lenw-1}这些位置开始检查，才能不遗漏
	for initialIndex := 0; initialIndex < lenw; initialIndex++ {
		index := initialIndex
		// moveIndex 让 index 指向下一个单词
		moveIndex := func() {
			// index 指向下一个单词的同时，需要同时修改统计记录

			// 增加 index 指向的 word 可出现次数一次，
			remainNum[s[index:index+lenw]]++
			// 连续符合条件的单词数减少一个
			count--
			// index 后移一个 word 的长度
			index += lenw
		} // 把这个匿名函数放在这里，可以避免繁琐的函数参数

		initRecord()

		for index+numWords*lenw <= lens { // 注意，这里是有等号的
			word := s[index+count*lenw : index+(count+1)*lenw]
			remainTimes, ok := remainNum[word]

			switch {
			case !ok:
				// 出现了不在 words 中的单词
				// 从 word 后面一个单词，重新开始统计
				index += lenw * (count + 1)
				if count != 0 {
					// 统计记录已经被修改，需要再次初始化
					initRecord()
				}
			case remainTimes == 0:
				// word 的出现次数上次就用完了
				// 说明s[index:index+(count+1)*lenw]中有单词多出现了
				moveIndex()
				// 这个case会连续出现
				// 直到s[index:index+(count+1)*lenw]中所有单词的出现次数都不超标
				//
				// 在 moveIndex() 的过程中，index+(count+1)*lenw 保持值不变
			default:
				// ok && remainTimes > 0，word 符合出现的条件

				// 更新统计记录
				remainNum[word]--
				count++

				// 检查 words 能否排列组合成 s[index:index+count*lenw]
				if count == numWords {
					res = append(res, index)

					// 把 index 指向下一个单词
					// 开始下一个统计
					moveIndex()
				}
			}
		}
	}

	return res
}
