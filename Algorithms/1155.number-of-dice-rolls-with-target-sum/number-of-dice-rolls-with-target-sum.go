package problem1155

const mod = 1e9 + 7

func numRollsToTarget(d int, f int, target int) int {
	rec := [1001]int{}

	var dp func(int, int) int
	dp = func(d, target int) int {
		if d == 0 && target == 0 {
			return 1
		}
		if d == 0 || target < 0 {
			return 0
		}
		if rec[target] > 0 {
			return rec[target]
		}
		res := 0
		for i := 1; i <= f; i++ {
			res += dp(d-1, target-i)
		}
		res %= mod
		rec[target] = res
		return res
	}

	return dp(d, target)
}
