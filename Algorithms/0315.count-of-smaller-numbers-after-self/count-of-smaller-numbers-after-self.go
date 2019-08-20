package problem0315

// ref: https://leetcode.com/problems/count-of-smaller-numbers-after-self/discuss/76584/Mergesort-solution
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
	merge := func(left, right []entry) []entry {
		res := make([]entry, 0, len(left)+len(right))
		var pop entry
		for len(left) > 0 && len(right) > 0 {
			if left[0].num > right[0].num {
				pop, left = left[0], left[1:]
				// for any i
				// pop.num > right[i].num, and
				// pop.index < right[i].index
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

	var sort func([]entry) []entry
	sort = func(es []entry) []entry {
		size := len(es)
		if size < 2 {
			return es
		}
		mid := size / 2
		return merge(sort(es[:mid]), sort(es[mid:]))
	}

	sort(enum)

	return count
}
