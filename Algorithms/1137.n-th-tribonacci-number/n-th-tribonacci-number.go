package problem1137

func tribonacci(n int) int {
	mem := make([]int, n+1)
	mem[0] = 0
	mem[1] = 1
	mem[2] = 1

	var t func(int) int
	t = func(n int) int {
		if n == 0 || mem[n] > 0 {
			return mem[n]
		}
		res := t(n-1) + t(n-2) + t(n-3)
		mem[n] = res
		return res
	}

	return t(n)
}
