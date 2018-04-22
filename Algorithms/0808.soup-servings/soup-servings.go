package problem0808

var record = map[[2]int]float64{}

func soupServings(N int) float64 {
	if N > 20000 {
		return 1
	}

	res := serve(N, N)

	return float64(int(res*1E5)) / 1E5
}

func serve(a, b int) float64 {
	if res, ok := record[[2]int{a, b}]; ok {
		return res
	}

	if a <= 0 && b > 0 {
		return 1
	}

	if a <= 0 && b <= 0 {
		return 0.5
	}

	if b <= 0 && a > 0 {
		return 0
	}

	res := 0.25*serve(a-100, b) +
		0.25*serve(a-75, b-25) +
		0.25*serve(a-50, b-50) +
		0.25*serve(a-25, b-75)

	record[[2]int{a, b}] = res

	return res
}
