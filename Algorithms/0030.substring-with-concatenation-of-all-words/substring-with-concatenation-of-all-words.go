package Problem0030

func findSubstring(s string, words []string) []int {
	res := []int{}
	if len(s) == 0 || len(words) == 0 || len(s) < len(words)*len(words[0]) {
		return res
	}

	initwMaps := func(wMaps map[string]int, words []string) {
		for _, word := range words {
			wMaps[word]++
		}
	}

	wMaps := make(map[string]int, len(words))
	lens, lenws, lenw := len(s), len(words), len(words[0])

	for slide := 0; slide < lenw; slide++ {
		initwMaps(wMaps, words)
		i, count := 0, 0
		for i <= lens-slide-lenws*lenw {
			tmp := s[i+slide+count*lenw : i+slide+(count+1)*lenw]
			if num, ok := wMaps[tmp]; ok && num != 0 {
				wMaps[tmp]--
				count++
				if count == lenws {
					res = append(res, i+slide)
					wMaps[s[i+slide:i+slide+lenw]]++
					count--
					i += lenw
				}
			} else if ok {
				wMaps[s[i+slide:i+slide+lenw]]++
				count--
				i += lenw

			} else {
				if count != 0 {
					initwMaps(wMaps, words)
				}
				i += lenw * (count + 1)
				count = 0
			}
		}
	}

	return res
}
