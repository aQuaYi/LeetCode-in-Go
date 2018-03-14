package problem0517

func findMinMoves(machines []int) int {
	ok, avg := isOK(machines)
	if !ok {
		return -1
	}

	cnt := 0
	res := 0

	for _, m := range machines {
		cnt += m - avg
		res = max(max(res, abs(cnt)), m-avg)
	}
	// 解题思路
	// https://leetcode.com/problems/super-washing-machines/discuss/

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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
