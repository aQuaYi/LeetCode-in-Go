package problem1109

func corpFlightBookings(bookings [][]int, n int) []int {
	res := make([]int, n)
	for _, b := range bookings {
		i, j, k := b[0], b[1], b[2]
		for m := i - 1; m < j; m++ {
			res[m] += k
		}
	}

	return res
}
