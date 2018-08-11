package problem0871

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	size := len(stations)
	dp := make([]int, size+1)

	f := startFuel
	for i, s := range stations {
		if f < s[0] {
			return -1
		}
		f = f - s[0] + s[1]
		dp[i] = i + 1
	}
	dp[size-1] = size

	fs := make([]int, size+1)
	miles := 0
	for i := range fs {
		miles += stations[i][0]
		fs[i] = startFuel - miles
	}
	if fs[size] >= 0 {
		return 0
	}

	stops := 0

	for i := 0; i < size; i++ {

	}

	return 0
}
