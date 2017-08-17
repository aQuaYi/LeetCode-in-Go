package Problem0605

func canPlaceFlowers(bed []int, n int) bool {
	length := len(bed)

	if n == 0 {
		return true
	}

	if length == 1 {
		if bed[0] == 0 && n == 1 {
			return true
		}
		return false
	}

	if length == 2 {
		if bed[0] == 0 && bed[1] == 0 && n == 1 {
			return true
		}
		return false
	}

	// 处理 length >= 3 的情况
	// 处理头部
	i := 0
	if bed[0] == 0 && bed[1] == 0 {
		bed[0] = 1
		n--
		if n == 0 {
			return true
		}
	}
	// 处理中间部分
	for i < length {
		for i+1 < length && bed[i+1] != 0 {
			i++
		}

		if i+3 < length && bed[i+2] == 0 && bed[i+3] == 0 {
			bed[i+2] = 1
			i += 2
			n--
			if n == 0 {
				return true
			}
		} else {
			i++
		}
	}
	// 处理尾部
	if bed[length-1] == 0 && bed[length-2] == 0 {
		n--
	}

	return n < 1
}
