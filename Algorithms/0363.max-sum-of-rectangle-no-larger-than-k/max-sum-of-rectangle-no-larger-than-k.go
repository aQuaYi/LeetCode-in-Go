package problem0363

func maxSumSubmatrix(mat [][]int, target int) int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return 0
	}

	m, n := len(mat), len(mat[0])
	// 本算法的复杂度是 O(M*M*N*logN)
	// 让 N = max(m,n) 可以加快程序
	M := min(m, n)
	N := max(m, n)
	// 从此开始称， M 为行，N 为列

	ans := -1 << 63

	var findTarget func([]int, int, int) int
	// 运用归并排序，在给 sums 排序的同时
	// 查找符合 <=target 的最大值
	findTarget = func(sums []int, beg, end int) int {
		if beg+1 >= end {
			return -1 << 63
		}

		mid := (beg + end) >> 1

		res := max(findTarget(sums, beg, mid), findTarget(sums, mid, end))
		if res == target {
			return res
		}

		l, r := beg, mid
		// sums[mid:end] 中任意的的 sums[r] 和
		// sums[beg:mid] 中任意的的 sums[l] 相减
		// temp = sums[r] - sums[l]都是某个子矩阵的所有元素之和
		// 此时，由于 sums[beg:mid] 和 sums[mid:end] 都是递增的
		// 可以避免不必要的检查
		for l < mid && r < end {
			temp := sums[r] - sums[l]
			switch {
			case temp < target:
				res = max(res, temp)
				// r++ 可以让 temp 变大
				// 这样才可能找到 更大的 res
				r++
			case temp > target:
				// temp > target
				// l++ 能让 temp 变小
				l++
			default:
				// 命中 target
				// 结束程序
				return target
			}
		}

		// 另外使用归并排序，这样的程序更简洁
		copy(sums[beg:end], merge(sums[beg:mid], sums[mid:end]))

		return res
	}

	var iFirst, iLast, j, minSum, maxSub int
	var temp, sums []int

	for iFirst = 0; iFirst < M; iFirst++ {
		// temp[j] 表示 mat[iFirst:iLast+1][j] 中所有元素之和
		temp = make([]int, N)
		for iLast = iFirst; iLast < M; iLast++ {
			// sums[j] 是 mat[iFirst:iLast+1][:j+1] 中所有元素之和
			// sums[j2] - sums[j1], j1 < j2 表示
			// mat[iFisrt:iLast+1][j1:j2+1] 中所有元素之和
			// sums 中添加 0 为了 sums[j] - 0  , 表示
			// mat[iFisrt:iLast+1][:j+1] 中所有元素之和
			sums = []int{0}
			// maxSub 是 mat[iFirst:iLast+1][:] 中所有子矩阵中，所有元素之和的 最大值
			maxSub = -1 << 63
			// minSum 是 sums 中的 最小值
			minSum = 0

			for j = 0; j < N; j++ {
				// 分情况更新 temp[j]
				if m < n {
					temp[j] += mat[iLast][j]
				} else {
					temp[j] += mat[j][iLast]
				}

				sums = append(sums, sums[len(sums)-1]+temp[j])
				// 求解 maxSub 的算法解释，看这里
				// https://www.youtube.com/watch?v=yCQN096CwWM
				maxSub = max(maxSub, sums[len(sums)-1]-minSum)
				minSum = min(minSum, sums[len(sums)-1])
				// TODO: 弄清楚这里的内容
			}

			// ans < target
			// ---- ans ---- target --->
			// 分 3 中情况讨论 maxSub 的值
			switch {
			case maxSub < target:
				// 可知，此时，maxSub == findTarget(sums,0,N+1)
				// 这时就体现了 maxSub 的巨大作用了
				// 省掉了运行 findTarget 的时间
				ans = max(ans, maxSub)
			case maxSub > target:
				// mat[iFirst:iLast+1][:] 中可能有个子矩阵的所有元素之和，
				// 比 ans 更接近 target
				// 需要运行 findTarget 查找
				tempAns := findTarget(sums, 0, N+1)
				if tempAns == target {
					// 找到答案，可以结束程序了
					return target
				}
				ans = max(ans, tempAns)
			default:
				// 找到答案，可以结束程序了
				return target
			}
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// a, b 都是升序的切片
// merge 把 a，b 合并起来，并保持升序
func merge(a, b []int) []int {
	lenA, lenB := len(a), len(b)
	temp := make([]int, lenA+lenB)

	var i, j, k int
	for i < lenA && j < lenB {
		if a[i] < b[j] {
			temp[k] = a[i]
			i++
		} else {
			temp[k] = b[j]
			j++
		}
		k++
	}

	if i == lenA {
		copy(temp[k:], b[j:])
	} else {
		copy(temp[k:], a[i:])
	}

	return temp
}
