package problem0202

func isHappy(n int) bool {
	slow, fast := n, trans(n)
	for slow != fast {
		slow = trans(slow)
		fast = trans(trans(fast))
	}
	if slow == 1 {
		return true
	}
	return false
}

func trans(n int) int {
	res := 0
	for n != 0 {
		res += (n % 10) * (n % 10)
		n /= 10
	}
	return res
}
