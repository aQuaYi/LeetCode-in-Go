package problem0995

func minKBitFlips(A []int, K int) int {
	size := len(A)
	res := 0
	for i := 0; i < size; i++ {
		if A[i] == 1 {
			continue
		}
		if i+K > size {
			return -1
		}
		res++
		for j := i; j < i+K; j++ {
			A[j] ^= 1 // flip
		}
	}

	return res
}
