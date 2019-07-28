package problem1109

func corpFlightBookings(bookings [][]int, n int) []int {
	res := make([]int, n+1)
	for _, b := range bookings {
		i, j, k := b[0], b[1], b[2]
		res[i-1] += k
		res[j] -= k
	}

	for i := 1; i < n; i++ {
		res[i] += res[i-1]
	}

	return res[:n]
}
