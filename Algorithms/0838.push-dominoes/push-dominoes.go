package problem0838

func pushDominoes(dominoes string) string {
	return oneSecond([]byte(dominoes))
}

func oneSecond(d []byte) string {
	old := string(d)
	size := len(d)

	for i := 0; i+1 < size; i++ {
		if d[i] != '.' {
			continue
		}
		if d[i+1] == 'L' &&
			(i == 0 || old[i-1] != 'R') {
			d[i] = 'L'
		}
	}

	for i := size - 1; 0 < i; i-- {
		if d[i] != '.' {
			continue
		}
		if d[i-1] == 'R' &&
			(i+1 == size || old[i+1] != 'L') {
			d[i] = 'R'
		}
	}

	newState := string(d)

	if newState == old {
		return newState
	}

	return oneSecond(d)
}
