package problem0777

import (
	"strings"
)

func canTransform(start string, end string) bool {
	// L 可以往左移
	// R 可以往右移
	// 但是，
	// L 没法移动到 R 的左边
	// R 没法移动到 L 的右边
	// 所有，L 和 R 互相为对方划清了移动范围
	// 无论 L 和 R 怎么移动，消去 X 后，start 和 end 中剩下的部分，应该是一样的。
	return strings.Replace(start, "X", "", -1) == strings.Replace(end, "X", "", -1) &&
		// 由于 L 不能往右移动，所以，每一个 L 的索引号只能变小
		isOK(idxs(start, 'L'), idxs(end, 'L'), isMoreOrEqual) &&
		// 由于 R 不能往左移动，所以，每一个 R 的索引号只能变大
		isOK(idxs(start, 'R'), idxs(end, 'R'), isLessOrEqual)
	// 换句话说，
	// 以上判断都是在检查 end 中的 L 和 R 的索引值变化有没有超出范围
	// 消去，检查了 end 中每个 L 的下限和每个 R 的上限
	// 然后
	// 检查了 end 中每个 L 的上限和每个 R 的下限
}

// idxs 返回 s 中所有 b 字符的索引号
func idxs(s string, b byte) []int {
	res := make([]int, 0, len(s))
	for i := range s {
		if s[i] == b {
			res = append(res, i)
		}
	}
	return res
}

// 对于任意的 i ，都能使的 isAvailable(a[i], b[i]) 返回 true
// 则 isOK 返回 true
func isOK(a, b []int, isAvailable func(x, y int) bool) bool {
	for i := range a {
		if !isAvailable(a[i], b[i]) {
			return false
		}
	}
	return true
}

func isLessOrEqual(x, y int) bool {
	if x <= y {
		return true
	}
	return false
}

func isMoreOrEqual(x, y int) bool {
	if x >= y {
		return true
	}
	return false
}
