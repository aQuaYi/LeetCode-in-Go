package problem0780

// 考虑到 x y 的范围是 [1, 10^9]
// (x, y) 在变形的过程中，只会越变越大
// 那么，反过来看，想要得到 (tx,ty)
// 	当 tx > ty 时， 必有 (tx-ty,ty) -> (tx,ty)
// 	当 tx < ty 时， 必有 (tx,ty-tx) -> (tx,ty)
// 所以，通过递归地减小 (tx,ty) 的值，可以很快的得到答案
func reachingPoints(sx int, sy int, tx int, ty int) bool {
	// 让 tx 比 ty 大
	if tx < ty {
		sx, sy = sy, sx
		tx, ty = ty, tx
	}

	var helper func(int, int) bool
	helper = func(x, y int) bool {
		// 找到答案
		if sx == x && sy == y {
			return true
		}

		// 确定无法找到答案
		if sx > x || sy > y ||
			x == y {
			// x == y 也返回 false，是因为 x-y == 0，超过了取值范围
			return false
		}

		if sy == y {
			// 当 sy == y 时
			// 只有 x-y 这一种运算了。
			// 就可以依据 sx%sy == x%y 判断结果
			return sx%sy == x%y
		}

		// 根据对称性
		// 直接把较小的数，放在 y 的位置
		return helper(y, x%y)
	}

	return helper(tx, ty)
}
