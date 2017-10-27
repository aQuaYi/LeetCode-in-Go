package Problem0335

func isSelfCrossing(x []int) bool {
	var n, w, s, e int
	for i := range x {
		switch i % 4 {
		case 0:
			n += x[i]
		case 1:
			w += x[i]
		case 2:
			s += x[i]
		default:
			e += x[i]
		}
	}

	return (n-s >= 0) && (e-w >= 0)
}
