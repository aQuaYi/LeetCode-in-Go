package problem1015

func smallestRepunitDivByK(K int) int {
	if K%2 == 0 {
		return -1
	}
	res := 1
	N := 1
	for N%K != 0 {
		N = N*10 + 1
		res++
	}
	return res
}
