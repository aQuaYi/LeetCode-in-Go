package problem1017

func baseNeg2(N int) string {
	if N == 0 {
		return "0"
	}

	B := make([]byte, 0, 30)
	for N > 0 {
		switch N & 3 {
		case 0, 1:
			B = append(B, byte(N&1)+'0', '0')
		default: // 2,3
			B = append(B, byte(N&1)+'0', '1')
			N += 4
			// 2^(2n)-2^(2n-1) == 2^(2n-1)
			// 所以，把 N 转换成二进制后，每在奇数位上看见 1 就在其左边的偶数位上加 1
			// 就可以转换成以 -2 为基的二进制了
		}
		N >>= 2
	}

	swap(B)

	if B[0] == '0' { // no lead 0 except for N is 0
		B = B[1:]
	}

	return string(B)
}

func swap(B []byte) {
	i, j := 0, len(B)-1
	for i < j {
		B[i], B[j] = B[j], B[i]
		i++
		j--
	}
}
