package problem0826

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	res := 0
	for _, w := range worker {
		p := 0
		for i := range difficulty {
			if difficulty[i] > w {
				continue
			}
			if p < profit[i] {
				p = profit[i]
			}
		}
		res += p
	}
	return res
}
