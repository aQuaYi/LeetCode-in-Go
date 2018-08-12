package problem0860

func lemonadeChange(bills []int) bool {
	changes := [5]int{}

	for _, b := range bills {
		changes[b/5]++

		if b == 5 {
			continue
		}

		b -= 5

		for i := 4; i > 0; i-- {
			if i*5 > b {
				continue
			}

			for changes[i] > 0 && b-i*5 >= 0 {
				b -= i * 5
				changes[i]--
			}
		}

		if b > 0 {
			return false
		}
	}

	return true
}
