package problem0605

func canPlaceFlowers(bed []int, n int) bool {
	l := len(bed)
	for i := 0; i < l; i++ {
		if bed[i] == 0 && // 检查 i 的值
			((i+1 < l && bed[i+1] == 0) || i+1 >= l) && // 检查 i+1 的值
			((i-1 >= 0 && bed[i-1] == 0) || i-1 < 0) { // 检查 i-1 的值
			bed[i] = 1
			n--
			if n <= 0 {
				return true
			}
		}
	}

	return n <= 0
}
