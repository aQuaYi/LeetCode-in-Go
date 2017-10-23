package Problem0354

import "sort"

func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) <= 1 {
		return len(envelopes)
	}

	sort.Sort(sortedEnvelopes(envelopes))

	var tails []int

	for i := 0; i < len(envelopes); i++ {
		lo, hi := 0, len(tails)
		for lo < hi {
			mid := (lo + hi) / 2
			if envelopes[i][1] <= tails[mid] {
				hi = mid
			} else {
				lo = mid + 1
			}
		}

		if lo == len(tails) {
			tails = append(tails, envelopes[i][1])
		} else {
			tails[lo] = envelopes[i][1]
		}
	}

	return len(tails)
}

type sortedEnvelopes [][]int

func (s sortedEnvelopes) Len() int {
	return len(s)
}

func (s sortedEnvelopes) Less(i, j int) bool {
	if s[i][0] != s[j][0] {
		return s[i][0] < s[j][0]
	}
	return s[i][1] > s[j][1]
}

func (s sortedEnvelopes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
