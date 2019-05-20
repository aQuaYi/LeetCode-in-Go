package problem0974

func subarraysDivByK(A []int, K int) int {
	count := [10000]int{}
	count[0] = 1
	prefix, res := 0, 0
	for _, a := range A {
		prefix = (prefix + a%K + K) % K
		res += count[prefix] // 减去相同的 prefix，剩下的就都是 K 的倍数
		count[prefix]++
	}
	return res
}
