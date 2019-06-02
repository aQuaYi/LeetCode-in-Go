package problem0995

func minKBitFlips(A []int, K int) int {
	size := len(A)

	swap := make([]int, size+1)
	flag := 1

	res := 0
	for i := 0; i < size; i++ {
		flag ^= swap[i]
		if A[i] == flag {
			continue
		}
		if i+K > size {
			return -1
		}
		res++
		flag ^= 1     // swap flag
		swap[i+K] = 1 // swap flag back when i+K
	}

	return res
}
