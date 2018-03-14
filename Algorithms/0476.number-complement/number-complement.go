package problem0476

func findComplement(num int) int {
	temp := num
	res := 0
	for temp > 0 {
		temp >>= 1
		res <<= 1
		res++
	}

	return res ^ num
}
