package problem0932

// ref: https://leetcode.com/problems/beautiful-array/discuss/186679/C%2B%2BJavaPython-Odd-%2B-Even-Pattern-O(N)
//

func beautifulArray(N int) []int {
	res := []int{1}
	for len(res) < N {
		tmp := make([]int, 0, len(res)*2)
		for _, v := range res {
			if v*2-1 <= N {
				tmp = append(tmp, v*2-1)
			}
		}
		for _, v := range res {
			if v*2 <= N {
				tmp = append(tmp, v*2)
			}
		}
		res = tmp
	}

	return res
}

/**
 * 以 N == 10 为例，res 的变化过程是
 * [1]
 * [1 2]
 * [1 3 2 4]
 * [1 5 3 7  2  6 4  8]
 * [1 9 5 13 3 11 7 15 2 10 6 14 4 12 8 16]
 * 长度超过 N，进行裁剪，删除 >N=10 的数字。裁剪不改变 beautiful
 * [1 9 5    3    7    2 10 6    4    8]
 *
 * 每个下一行的
 *     左半部分，都是 上一行×2-1
 *     右半部分，都是 上一行×2
 * 由文章中证明，× 和 - 不会改变 beautiful
 * 所以，左右两边都是 beautiful 的
 *
 * 左右合并的时候，只需要讨论 Ai 和 Aj 分别位于左右两边的情况。
 * 左边的 Ai 是奇数，右边的 Aj 是偶数。所以 Ai+Aj 是奇数。
 * 而 Ak×2 是偶数，所以，不存在 Ak×2 == Ai+Aj
 * A 仍然是 beautiful 的
 */
