package problem1017

func baseNeg2(N int) string {
	B := make([]byte, 0, 30)
	for N > 0 {
		switch N & 3 {
		case 0:
			B = append(B, '0', '0')
		case 1:
			B = append(B, '1', '0')
		case 2:
			B = append(B, '0', '1')
			N += 4
		case 3:
			B = append(B, '1', '1')
			N += 4
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
