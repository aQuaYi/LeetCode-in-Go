package problem0878

func nthMagicalNumber(N int, A int, B int) int {
	if A > B {
		A, B = B, A
	}

	num := A - 1
	count := 0
	for count < N {
		num++
		if num%A == 0 || num%B == 0 {
			count++
		}
	}

	return num
}
