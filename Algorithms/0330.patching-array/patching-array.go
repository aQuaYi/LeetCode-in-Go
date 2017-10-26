package Problem0330

import "sort"

func minPatches(nums []int, n int) int {
	var i, idx, size int

	idxLen, count := log2(n)+1, 0

	pow2 := make([]int, 1, idxLen)
	pow2[0] = 1
	for i = 1; i < idxLen; i++ {
		pow2[i] = pow2[i-1] * 2
	}

	sort.Ints(nums)

	i, idx, size = 0, 0, len(nums)
	for i < size && idx < idxLen {
		if nums[i] > pow2[idx] {
			count++
			idx++
			continue
		}

		if nums[i] == pow2[idx] {
			idx++
			i++
			continue
		}

	}

	return count
}

// log2(n) 返回 n 以 2 为底的对数的整数部分
func log2(n int) int {
	res := 0
	for n > 1 {
		n = n >> 1
		res++
	}
	return res
}
