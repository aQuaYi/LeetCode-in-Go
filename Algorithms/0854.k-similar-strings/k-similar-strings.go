package problem0854

func kSimilarity(a, b string) int {
	if a == b {
		return 0
	}

	hasSeen := make(map[string]bool, 4096)
	hasSeen[a] = true
	queue := make([]string, 1, 4096)
	queue[0] = a

	strSize := len(a)
	res := 0

	// BFS
	// 看见 shortest 就要想到 BFS
	for {
		res++
		for countDown := len(queue); countDown > 0; countDown-- {
			// 从 queue 中取出一个候选字符串
			s := queue[0]
			queue = queue[1:]

			// 跳过相等的字符
			i := 0
			for s[i] == b[i] {
				i++
			}
			for j := i + 1; j < strSize; j++ {
				if s[j] == b[j] || // 依然跳过相等的字符
					// 还要保证 s 中的 i，j 互换后，s[j] == b[j]
					// 要不然，就是没有意义的交换
					s[i] != b[j] {
					continue
				}

				temp := swap(s, i, j)
				if temp == b {
					return res
				}

				if !hasSeen[temp] {
					hasSeen[temp] = true
					queue = append(queue, temp)
				}
			}
		}
	}
}

func swap(s string, i, j int) string {
	bs := []byte(s)
	bs[i], bs[j] = bs[j], bs[i]
	return string(bs)
}
