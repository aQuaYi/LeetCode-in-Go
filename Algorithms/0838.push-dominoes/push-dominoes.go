package problem0838

func pushDominoes(dominoes string) string {
	d := "L" + dominoes + "R"
	size := len(d)
	res := make([]byte, 0, size)

	for i, j := 0, 1; j < size; j++ {
		if d[j] == '.' {
			continue
		}

		if 0 < i {
			res = append(res, d[i])
		}

		mid := j - i - 1

		switch {
		case d[i] == 'R' && d[j] == 'L':
			// 'R......L' => 'RRRRLLLL' or 'RRRR.LLLL'
			for k := 0; k < mid/2; k++ {
				res = append(res, 'R')
			}
			if mid%2 == 1 {
				res = append(res, '.')
			}
			for k := 0; k < mid/2; k++ {
				res = append(res, 'L')
			}
		case d[i] == 'L' && d[j] == 'R':
			// 'L......R' => 'L......R'
			for k := 0; k < mid; k++ {
				res = append(res, '.')
			}
		default:
			//  'R......R' => 'RRRRRRRR'
			//  'L......L' => 'LLLLLLLL'
			for k := 0; k < mid; k++ {
				res = append(res, d[i])
			}
		}

		i = j
	}

	return string(res)
}
