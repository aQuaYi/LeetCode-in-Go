package problem1015

// ref: https://leetcode.com/problems/smallest-integer-divisible-by-k/discuss/260852/JavaC%2B%2BPython-O(1)-Space-with-Proves-of-Pigeon-Holes
func smallestRepunitDivByK(K int) int {
	if K%2 == 0 || K%5 == 0 {
		return -1
	}

	r := 0
	for N := 1; N <= K; N++ {
		r = (r*10 + 1) % K
		if r == 0 {
			return N
		}
	}

	return -1
}
