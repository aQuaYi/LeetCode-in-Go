package problem0621

func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}

	rec := [26]int{}
	for _, c := range tasks {
		rec[c-'A']++
	}

	// 找到最多的任务数
	most := 0
	for _, num := range rec {
		if most < num {
			most = num
		}
	}

	// 整个过程，至少有 idles 个空档需要被填充
	idles := (most - 1) * (n + 1)

	// 使用每一个 task 去填充 idles
	for _, num := range rec {
		// 可以假想 idles+1 个空档，被 most 的 task 等间距n填充后，
		// 分割成了 most-1 个包含连续 n 个空档的区间
		// 然后，同样的任务，只能往每个区间中去填充。
		idles -= min(most-1, num)
		// 由于是使用 most 做的分割
		// 如果最后的 idles < 0 只能说明
		// len(tasks) > (most - 1) * (n + 1)+1
	}

	return len(tasks) + max(0, idles)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
