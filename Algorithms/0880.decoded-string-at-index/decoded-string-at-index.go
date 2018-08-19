package problem0880

func decodeAtIndex(S string, K int) string {
	N, i := 0, 0
	b := byte(0)

	for i = 0; N < K; i++ {
		b = S[i]
		if isDigit(b) {
			N *= int(b - '0')
		} else {
			N++
		}
	}

	for {
		i--
		b = S[i]
		if isDigit(b) {
			N /= int(b - '0')
			K %= N
		} else {
			if K == 0 || K == N {
				return string(b)
			}
			N--
		}
	}

}

func isDigit(b byte) bool {
	return b <= '9'
}
