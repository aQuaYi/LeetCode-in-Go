package problem0880

func decodeAtIndex(S string, K int) string {
	/** lenSi 是 S[:i] 展开后的长度。*/
	lenSi, i := 0, 0

	for ; lenSi < K; i++ {
		char := S[i]
		if isDigit(char) {
			lenSi *= int(char - '0')
		} else {
			lenSi++
		}
	}

	for {
		i--
		char := S[i]
		if isDigit(char) {
			lenSi /= int(char - '0')
			K %= lenSi
		} else {
			if K == 0 || // "leet3code2" K = 8 会导致 K == 0 成立
				K == lenSi {
				return string(char)
			}
			lenSi--
		}
	}

}

func isDigit(b byte) bool {
	return b <= '9'
}
