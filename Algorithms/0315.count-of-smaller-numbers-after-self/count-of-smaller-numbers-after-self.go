package problem0315

type entry struct {
	num, index int
}

func countSmaller(nums []int) []int {
	n := len(nums)
	enum := make([]entry, n)
	for i, n := range nums {
		enum[i] = entry{num: n, index: i}
	}

	count := make([]int, n)

	var sort func([]entry) []entry
	var merge func([]entry, []entry) []entry

	sort = func(es []entry) []entry {
		size := len(es)
		if size < 2 {
			return es
		}
		mid := size / 2
		return merge(sort(es[:mid]), sort(es[mid:]))
	}

	merge = func(left, right []entry) []entry {
		m, n := len(left), len(right)
		res := make([]entry, 0, m+n)
		var pop entry
		for len(left) > 0 && len(right) > 0 {
			if left[0].num > right[0].num {
				pop, left = left[0], left[1:]
				count[pop.index] += len(right)
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
