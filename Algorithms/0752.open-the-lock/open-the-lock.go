package problem0752

func openLock(deadends []string, target string) int {
	isDeadends := dealDeadends(deadends)
	return bfs(convert(target), isDeadends)
}

func bfs(target [4]int, isDeadends map[[4]int]bool) int {
	queue := [][4]int{convert("0000")}
	if isDeadends[queue[0]] {
		return -1
	}
	if target == queue[0] {
		return 0
	}
	isChecked := make(map[[4]int]bool, 1e4)
	isChecked[queue[0]] = true
	leng := 1
	count := 1

	for len(queue) > 0 {
		if count == 0 {
			leng++
			count = len(queue)
		}

		for i := 0; i < 4; i++ {
			a := queue[0]
			a[i] = (a[i] + 9) % 10
			if !isDeadends[a] && !isChecked[a] {
				if a == target {
					return leng
				}
				queue = append(queue, a)
				isChecked[a] = true
			}
			b := queue[0]
			b[i] = (b[i] + 1) % 10
			if !isDeadends[b] && !isChecked[b] {
				if b == target {
					return leng
				}
				queue = append(queue, b)
				isChecked[b] = true
			}
		}

		queue = queue[1:]
		count--
	}

	return -1
}

func dealDeadends(deadends []string) map[[4]int]bool {
	res := make(map[[4]int]bool, len(deadends))
	for _, d := range deadends {
		res[convert(d)] = true
	}
	return res
}

func convert(s string) [4]int {
	res := [4]int{}
	for i := range res {
		res[i] = int(s[i] - '0')
	}
	return res
}
