package Problem0446

func numberOfArithmeticSlices(a []int) int {
	size := len(a)
	if size < 3 {
		return 0
	}

	res := 0

	var dfs func(int, int)
	dfs = func(idx, diff int) {
		if idx == size {
			return
		}
		for i := idx + 1; i < size; i++ {
			if a[i]-a[idx] == diff {
				res++
				dfs(i, diff)
			}
		}
	}

	for i := 0; i < size-2; i++ {
		for j := i + 1; j < size-1; j++ {
			dfs(j, a[j]-a[i])
		}

	}

	return res
}
