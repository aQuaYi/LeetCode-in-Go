package problem1010

func numPairsDivisibleBy60(time []int) int {
	rec := [60]int{}
	res := 0
	for _, t := range time {
		t %= 60
		res += rec[(60-t)%60]
		rec[t]++
	}
	return res
}
