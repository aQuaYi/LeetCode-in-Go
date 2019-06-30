package problem1041

func isRobotBounded(instructions string) bool {
	g, l, r := 0, 0, 0
	for _, b := range instructions {
		switch b {
		case 'G':
			g++
		case 'L':
			l++
		case 'R':
			r++
		}
	}
	return g == 0 || l != r
}
