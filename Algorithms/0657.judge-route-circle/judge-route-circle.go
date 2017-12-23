package Problem0657

var dx = map[byte]int{
	'R': 1,
	'L': -1,
	'U': 0,
	'D': 0,
}

var dy = map[byte]int{
	'R': 0,
	'L': 0,
	'U': -1,
	'D': 1,
}

func judgeCircle(moves string) bool {
	x, y := 0, 0
	for i := range moves {
		x += dx[moves[i]]
		y += dy[moves[i]]
	}

	return x == 0 && y == 0
}
