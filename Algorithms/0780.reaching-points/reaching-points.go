package problem0780

func reachingPoints(sx int, sy int, tx int, ty int) bool {
	if sx == tx && sy == ty {
		return true
	}

	if sx > tx || sy > ty ||
		tx == ty {
		return false
	}

	if tx < ty {
		sx, sy = sy, sx
		tx, ty = ty, tx
	}

	if sy == ty && sx%sy == tx%ty {
		return true
	}

	return reachingPoints(sy, sx, ty, tx%ty)
}
