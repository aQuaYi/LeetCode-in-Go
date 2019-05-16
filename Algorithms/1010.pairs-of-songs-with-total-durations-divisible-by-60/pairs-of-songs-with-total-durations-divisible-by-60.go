package problem1010

func numPairsDivisibleBy60(time []int) int {
	rec := [60]int{}
	for _, t := range time {
		rec[t%60]++
	}

	res := rec[0] * (rec[0] - 1) / 2
	res += rec[30] * (rec[30] - 1) / 2
	for i := 1; i < 30; i++ {
		res += rec[i] * rec[60-i]
	}

	return res
}
