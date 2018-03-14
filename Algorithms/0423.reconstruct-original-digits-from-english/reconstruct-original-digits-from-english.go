package problem0423

func originalDigits(s string) string {
	if len(s) == 0 {
		return ""
	}

	var counts [10]int

	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		case 'z':
			counts[0]++ // 0
		case 'w':
			counts[2]++ // 2
		case 'x':
			counts[6]++ // 6
		case 'u':
			counts[4]++ // 4
		case 'g':
			counts[8]++ // 8
		case 's':
			counts[7]++ // 7 & 6
		case 'v':
			counts[5]++ // 5 & 7
		case 'h':
			counts[3]++ // 3 & 8
		case 'i':
			counts[9]++ // 5 & 6 & 8 & 9
		case 'o':
			counts[1]++ // 1 & 2 & 4 & 0
		}
	}

	counts[3] -= counts[8]
	counts[7] -= counts[6]
	counts[5] -= counts[7]
	counts[1] -= counts[2] + counts[0] + counts[4]
	counts[9] -= counts[5] + counts[6] + counts[8]

	res := make([]byte, 0, len(s)/3)

	for b := byte(0); b < 10; b++ {
		for counts[b] > 0 {
			res = append(res, b+'0')
			counts[b]--
		}
	}

	return string(res)
}
