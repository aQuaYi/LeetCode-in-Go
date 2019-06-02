package problem0995

func minKBitFlips(A []int, K int) int {
	size := len(A)
	queue := make([]int, 0, size)
	flag := 1
	res := 0
	for i := 0; i < size; i++ {
		if len(queue) > 0 && i == queue[0] {
			flag ^= 1
			queue = queue[1:]
		}
		if A[i] == flag {
			continue
		}
		if i+K > size {
			return -1
		}
		res++
		flag ^= 1
		queue = append(queue, i+K)
	}

	return res
}
