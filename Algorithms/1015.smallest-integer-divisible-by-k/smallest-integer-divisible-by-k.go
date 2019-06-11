package problem1015

// ref: https://leetcode.com/problems/smallest-integer-divisible-by-k/discuss/260852/JavaC%2B%2BPython-O(1)-Space-with-Proves-of-Pigeon-Holes
func smallestRepunitDivByK(K int) int {
	if K%2 == 0 || K%5 == 0 {
		return -1
	}

	r, length := 1%K, 1
	for r != 0 && length <= K {
		r = (r*10 + 1) % K
		length++
	}

	return length
}
