package Problem0657

func judgeCircle(moves string) bool {
	bs := []byte(moves)
	var up, down, left, right int
	for i := range bs {
		switch bs[i] {
		case 'U':
			up++
		case 'D':
			down++
		case 'L':
			left++
		case 'R':
			right++
		}
	}

	return up == down && left == right
}
