package problem0860

func lemonadeChange(bills []int) bool {
	fives, tens := 0, 0

	for _, b := range bills {
		switch b {
		case 5:
			fives++
		case 10:
			fives--
			tens++
		case 20:
			if tens > 0 {
				// 找零的时候，尽量先给 10 元的整钱
				// 而不是两个 5 元
				tens--
				fives--
			} else {
				fives -= 3
			}
		}
		if fives < 0 || tens < 0 {
			return false
		}
	}

	return true
}
