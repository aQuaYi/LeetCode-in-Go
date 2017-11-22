package Problem0552

var m = 1000000007

func checkRecord(n int) int {
	if n == 1 {
		return 3
	}
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 0
	}

	res := checkRecord(n - 1)
	res += hasA(n - 1)
	res += afterL(n - 1)

	return res % m
}

func hasA(n int) int {
	if n == 1 {
		return 2
	}

	if n == 0 {
		return 1
	}

	if n < 0 {
		return 0
	}

	res := hasA(n-1) + // n 位为 “P”
		afterAL(n-1) // n 位为　“Ｌ”
	return res % m
}

func afterAL(n int) int {
	if n == 1 {
		return 2
	}
	if n <= 0 {
		return 0
	}
	res := hasA(n-1) + // n 位为 “P”
		hasA(n-2) // n 位为 “L”
	return res % m
}

func afterL(n int) int {
	if n == 1 {
		return 3
	}

	if n == 0 {
		return 1
	}

	if n < 0 {
		return 0
	}

	res := checkRecord(n - 1)
	res += hasA(n - 1)
	res += checkRecord(n-2) + hasA(n-2)

	return res % m
}
