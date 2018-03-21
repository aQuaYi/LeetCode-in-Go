package problem0780

func reachingPoints(sx int, sy int, tx int, ty int) bool {
	if sx == tx && sy == ty {
		return true
	}

	if sx > tx || sy > ty ||
		tx == ty {
		return false
	}

	if tx > ty {
		return reachingPoints(sx, sy, tx%ty, ty)
	}
	return reachingPoints(sx, sy, tx, ty%tx)
}
