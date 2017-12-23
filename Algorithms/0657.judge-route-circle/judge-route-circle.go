package Problem0657

func judgeCircle(moves string) bool {
	x, y := 0, 0
	for i := range moves {
		switch moves[i] {
		case 'R':
			x++
		case 'L':
			x--
		case 'U':
			y--
		case 'D':
			y++
		}
	}

	return x == 0 && y == 0
}
