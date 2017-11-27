package Problem0477

func totalHammingDistance(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			res += hd(nums[i], nums[j])
		}
	}

	return res
}

// 返回 a 和 b 的 hamming distance
func hd(a, b int) int {
	sum := 0
	t := a ^ b
	for t > 0 {
		if t&1 == 1 {
			sum++
		}
		t >>= 1
	}
	return sum
}
