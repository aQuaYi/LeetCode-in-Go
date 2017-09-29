package Problem0210

func findOrder(numCourses int, prerequisites [][]int) []int {
n , p := build(numCourses, prerequisites)
	return search(n, p)
}

func build(num int, requires [][]int) (next [][]int, pre []int) {
	next = make([][]int, num)
	pre = make([]int, num)

	for _, r := range requires {
		next[r[1]] = append(next[r[1]], r[0])
		pre[r[0]]++
	}

	return
}

func search(next [][]int, pre []int) []int {
	var i, j int
	n := len(pre)
	res := make([]int, n)
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			if pre[j] == 0 {
				break
			}
		}
		if j == n {
			return nil
		}

		pre[j] = -1
		res[i] = j
		for _, c := range next[j] {
			pre[c]--
		}
	}

	return res
}
