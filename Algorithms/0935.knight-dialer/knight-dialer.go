package problem0935

const mod = 1e9 + 7

var rec = [5000][10]int{}

var hop = [][]int{
	{4, 6},    // from 0
	{6, 8},    // from 1
	{7, 9},    // from 2
	{4, 8},    // from 3
	{3, 9, 0}, // from 4, 4 -> 7 -> * -> 0
	{},        // from 5
	{0, 1, 7}, // from 6
	{2, 6},    // from 7
	{1, 3},    // from 8
	{2, 4},    // from 9
}

func knightDialer(N int) int {
	if N == 1 {
		// now can chose 5
		return 10
	}

	res := 0
	for i := 0; i < 10; i++ {
		res += recur(N-1, i, &rec)
	}

	return res % mod
}

func recur(N, digit int, rec *[5000][10]int) int {
	r := (*rec)[N][digit]
	if r > 0 {
		return r
	}

	h := hop[digit]

	if N == 1 {
		l := len(h)
		(*rec)[N][digit] = l
		return l
	}

	res := 0
	for _, d := range h {
		res += recur(N-1, d, rec)
	}
	res %= mod

	(*rec)[N][digit] = res
	return res
}
