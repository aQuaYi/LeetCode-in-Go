package Problem0605

func canPlaceFlowers(bed []int, n int) bool {
	l := len(bed)
	for i := 0; i < l; i++ {
		// if 的判断语句分别检查了，i，i-1 和 i+1的情况
		if bed[i] == 0 &&
			((i+1 < l && bed[i+1] == 0) || i+1 >= l) &&
			((i-1 >= 0 && bed[i-1] == 0) || i-1 < 0) {
			bed[i] = 1
			n--
		}
	}

	return n <= 0
}
