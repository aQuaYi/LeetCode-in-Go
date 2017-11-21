package Problem0517

func findMinMoves(machines []int) int {
	res := 0

	ok, ave := isOK(machines)
	if !ok {
		return -1
	}

	for _, m := range machines {
		res += max(0, m-ave)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isOK(machines []int) (bool, int) {
	sum := 0
	for _, m := range machines {
		sum += m
	}
	return sum%len(machines) == 0, sum / len(machines)
}
