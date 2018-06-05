package problem0838

import (
	"bytes"
)

/**
 * 'R......R' => 'RRRRRRRR'
 * 'R......L' => 'RRRRLLLL' or 'RRRR.LLLL'
 * 'L......R' => 'L......R'
 * 'L......L' => 'LLLLLLLL'
 */
func pushDominoes(dominoes string) string {
	var res bytes.Buffer
	d := "L" + dominoes + "R"
	size := len(d)

	for i, j := 0, 1; j < size; j++ {
		if d[j] == '.' {
			continue
		}

		if 0 < i {
			res.WriteByte(d[i])
		}

		mid := j - i - 1

		switch {
		case d[i] == d[j]:
			for k := 0; k < mid; k++ {
				res.WriteByte(d[i])
			}
		case d[i] == 'R' && d[j] == 'L':
			for k := 0; k < mid/2; k++ {
				res.WriteByte('R')
			}
			if mid%2 == 1 {
				res.WriteByte('.')
			}
			for k := 0; k < mid/2; k++ {
				res.WriteByte('L')
			}
		default:
			for k := 0; k < mid; k++ {
				res.WriteByte('.')
			}
		}

		i = j
	}

	return res.String()
}
