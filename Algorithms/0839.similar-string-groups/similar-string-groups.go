package problem0839

func numSimilarGroups(A []string) int {
	size := len(A)
	count := size

	parent := make([]int, size)
	for i := 0; i < size; i++ {
		parent[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if x != parent[x] {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(x, y int) {
		x, y = find(x), find(y)
		if x != y {
			parent[y] = x
			count--
		}
	}

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if isSimilar(A[i], A[j]) {
				union(i, j)
			}
		}
	}

	return count
}

func isSimilar(a, b string) bool {
	hasSeen := [26]bool{}
	hasDouble := false
	count := 0
	for i := 0; i < len(a) && count < 3; i++ {
		if a[i] != b[i] {
			count++
			continue
		}
		if hasSeen[a[i]-'a'] {
			hasDouble = true
		}
		hasSeen[a[i]-'a'] = true
	}

	return count == 2 || (count == 0 && hasDouble)
}
