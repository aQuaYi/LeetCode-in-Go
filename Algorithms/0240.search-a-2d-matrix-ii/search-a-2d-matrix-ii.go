package Problem0240

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}

	return search(matrix, target, 0, m-1, 0, n-1)
}

func search(m [][]int, target, cl, cr, rl, rr int) bool {
	if cl > cr || rl > rr {
		return false
	}

	if target < m[cl][rl] || m[cr][rr] < target {
		return false
	}

	cm := (cl + cr) / 2
	rm := (rl + rr) / 2
	if m[cm][rm] == target {
		return true
	}

	if search(m, target, cl, cm-1, rm+1, rr) ||
		search(m, target, cm+1, cr, rl, rm-1) {
		return true
	}

	if target < m[cm][rm] {
		return search(m, target, cl, cm-1, rl, rm) || search(m, target, cm, cm, rl, rm-1)
	}

	return search(m, target, cm+1, cr, rm, rr) || search(m, target, cm, cm, rm+1, rr)
}
