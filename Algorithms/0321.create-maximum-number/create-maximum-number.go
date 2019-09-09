package problem0321

func maxNumber(A, B []int, k int) []int {
	m, n := len(A), len(B)
	res := make([]int, k)
	for i := max(0, k-n); i <= m && k-i >= 0; i++ {
		temp := combine(choose(A, i), choose(B, k-i))
		if isBigger(temp, res, 0, 0) {
			res = temp
		}
	}
	return res
}

func choose(A []int, k int) []int {
	stack, top := make([]int, k), -1
	n := len(A)
	for i := 0; i < n; i++ {
		for top >= 0 && stack[top] < A[i] &&
			top+n-i >= k {
			// 此时，stack 中 有 top+1 个元素
			// 需要保证消去 stack[top] 后的 top 个元素
			// 和 A[i:] 中的 n-i 个元素
			// 能够组成 k 个元素
			// 所以，top+(n-i)>=k
			top--
		}
		if top+1 < k {
			// 如果 stack 中的元素，还没有 k 个的话
			// 先用 A[i] 填上空缺
			top++
			stack[top] = A[i]
		}
	}
	return stack
}

// 混合 nums1 和 nums2 使得其组成的 []int 最大
func combine(A, B []int) []int {
	m, n := len(A), len(B)
	res := make([]int, 0, m+n)

	var i, j int
	for i < m && j < n {
		if isBigger(A, B, i, j) {
			res = append(res, A[i])
			i++
		} else {
			res = append(res, B[j])
			j++
		}
	}

	res = append(res, A[i:]...)
	res = append(res, B[j:]...)

	return res
}

func isBigger(A, B []int, i, j int) bool {
	m, n := len(A), len(B)
	for i < m && j < n {
		if A[i] > B[j] {
			return true
		} else if A[i] < B[j] {
			return false
		}
		i++
		j++
	}
	return m-i > n-j
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
