package problem0808

var results = [201][201]float64{}

func soupServings(N int) float64 {
	if N >= 5000 {
		return 1
	}
	return serve((N+24)/25, (N+24)/25)
}

func serve(a, b int) float64 {
	if a <= 0 && b > 0 {
		return 1
	}

	if a <= 0 && b <= 0 {
		return 0.5
	}

	if b <= 0 && a > 0 {
		return 0
	}

	if results[a][b] > 0 {
		return results[a][b]
	}

	results[a][b] = 0.25*serve(a-4, b) +
		0.25*serve(a-3, b-1) +
		0.25*serve(a-2, b-2) +
		0.25*serve(a-1, b-3)

	return results[a][b]
}
