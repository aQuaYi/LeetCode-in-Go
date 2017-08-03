package Problem0030

func findSubstring(s string, words []string) []int {
	res := []int{}
	if len(s) == 0 || len(words) == 0 || len(s) < len(words)*len(words[0]) {
		return res
	}
	// remainNum 用于记录 words 中单词还能出现的次数
	remainNum := make(map[string]int, len(words))
	lens, numWords, lenw := len(s), len(words), len(words[0])

	/*
		for slice := 0; slice < lenw; slice++ {
			begin := slice
			for begin+numWords*lenw <= lens {
				begin += lenw
			}
		}

		利用上面这个结构，可以不遗漏也不重复地检查所有可能的结果
		把 slice < lenw 改成 slice < lens 会出现 很多重复的结果
	*/
	for slice := 0; slice < lenw; slice++ {
		// 利用 slice 把s转换转换成了 lenw 个 字符串去检查
		// 例如， 当 slice == 1 时
		// 只用依次检查 s[1:1+lenw], s[1+lenw:1+lenw*2],...,s[1+lenw*(n):1+lenw*(n+1)],...这些单词是否条件
		initialise(remainNum, words)
		begin, count := slice, 0

		for begin+numWords*lenw <= lens {
			w := s[begin+count*lenw : begin+(count+1)*lenw]

			num, ok := remainNum[w]
			switch {
			case !ok:
				// 出现了words以外的单词
				// 从 w 后面一个字符，重新开始
				begin += lenw * (count + 1)
				if count != 0 {
					// wordMap 已经被修改，需要再次初始化
					initialise(remainNum, words)
					count = 0
				}
			case num == 0:
				// w 的出现次数超标了

				// 修改统计记录，需要把 begin 后移一个单词的长度
				// 增加 begin 处 word 可出现次数一次，
				remainNum[s[begin:begin+lenw]]++
				// 统计记录增加一次
				count--
				// begin 后移一个 word 的长度
				begin += lenw

				// 这个case会连续执行，直到begin指向上一个 w 的后面一个字符
			default:
				// ok && num != 0

				// 修改统计记录
				remainNum[w]--
				count++

				// 检查记录是否已经符合要求
				if count == numWords {
					res = append(res, begin)

					// 把begin指向下一个单词
					remainNum[s[begin:begin+lenw]]++
					count--
					begin += lenw
				}
			}
		}
	}
	return res
}

func initialise(wordMap map[string]int, words []string) {
	for _, word := range words {
		wordMap[word] = 0
	}
	for _, word := range words {
		wordMap[word]++
	}
}
