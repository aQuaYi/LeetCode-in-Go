package problem0825

func numFriendRequests(ages []int) int {
	size := len(ages)
	res := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == j || isNotRequest(ages[i], ages[j]) {
				continue
			}
			res++
		}
	}
	return res
}

func isNotRequest(a, b int) bool {
	if b <= a/2+7 ||
		b > a ||
		(b > 100 && a < 100) {
		return true
	}
	return false
}
