package problem0315

func countSmaller(nums []int) []int {
	n := len(nums)
	enum := make([][2]int, n)
	for i, n := range nums {
		enum[i] = [2]int{n, i}
	}

	count := make([]int, n)

	var sort func([][2]int) [][2]int
	var merge func([][2]int, [][2]int) [][2]int

	sort = func(es [][2]int) [][2]int {
		size := len(es)
		if size < 2 {
			return es
		}
		mid := size / 2
		return merge(sort(es[:mid]), sort(es[mid:]))
	}

	merge = func(left, right [][2]int) [][2]int {
		m, n := len(left), len(right)
		res := make([][2]int, 0, m+n)
		var pop [2]int
		for len(left) > 0 && len(right) > 0 {
			if left[0][0] > right[0][0] {
				pop, left = left[0], left[1:]
				count[pop[1]] += len(right)
			} else {
				pop, right = right[0], right[1:]
			}
			res = append(res, pop)
		}
		res = append(res, left...)
		res = append(res, right...)

		return res
	}

	sort(enum)

	return count
}
