package problem0287

func findDuplicate(a []int) int {
	slow, fast := a[0], a[a[0]]
	for slow != fast {
		slow, fast = a[slow], a[a[fast]]
	}

	slow = 0
	for slow != fast {
		slow, fast = a[slow], a[fast]
	}

	return slow
}
