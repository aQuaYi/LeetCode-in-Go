package problem0975

func oddEvenJumps(A []int) int {
	odds := make(map[int]bool, len(A))
	evens := make(map[int]bool, len(A))
	res := 0
	for i := 0; i < len(A); i++ {
		oddJump(A, i, odds, evens, &res)
	}
	return res
}

func oddJump(A []int, i int, odds, evens map[int]bool, res *int) bool {
	if len(A)-1 == i || odds[i] {
		*res++
		return true
	}
	if success, ok := odds[i]; ok && !success {
		return false
	}
	j := oddNumberedJumps(A, i)
	if j == -1 {
		odds[j] = false
		return false
	}
	odds[i] = evenJump(A, j, odds, evens, res)
	return odds[i]
}

func evenJump(A []int, i int, odds, evens map[int]bool, res *int) bool {
	if len(A)-1 == i || evens[i] {
		*res++
		return true
	}
	if success, ok := evens[i]; ok && !success {
		return false
	}
	j := evenNumberedJumps(A, i)
	if j == -1 {
		evens[j] = false
		return false
	}
	evens[i] = oddJump(A, j, odds, evens, res)
	return evens[i]
}

func oddNumberedJumps(A []int, i int) int {
	minNum := 100000
	index := -1
	for j := len(A) - 1; j > i; j-- {
		if A[i] <= A[j] && minNum >= A[j] {
			minNum = A[j]
			index = j
		}
	}
	return index
}

func evenNumberedJumps(A []int, i int) int {
	maxNum := -1
	index := -1
	for j := len(A) - 1; j > i; j-- {
		if A[i] >= A[j] && maxNum <= A[j] {
			maxNum = A[j]
			index = j
		}
	}
	return index
}
