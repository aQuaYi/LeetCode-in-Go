package problem1221

func balancedStringSplit(s string) int {
	res, count := 0, 0
	for _, b := range s {
		if b == 'L' {
			count++
		} else {
			count--
		}
		if count == 0 {
			res++
		}
	}
	return res
}
