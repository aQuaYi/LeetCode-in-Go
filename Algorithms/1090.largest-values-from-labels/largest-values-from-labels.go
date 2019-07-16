package problem1090

import "sort"

func largestValsFromLabels(values []int, labels []int, num_wanted int, use_limit int) int {
	for i, v := range values {
		values[i] = v<<16 + labels[i]
	}

	sort.Slice(values, func(i int, j int) bool {
		return values[i] > values[j]
	})

	count := [20001]int{}
	res := 0
	for _, v := range values {
		if count[v&0xFFFF] == use_limit {
			continue
		}
		res += v >> 16
		count[v&0xFFFF]++
		num_wanted--
		if num_wanted == 0 {
			break
		}
	}

	return res
}
