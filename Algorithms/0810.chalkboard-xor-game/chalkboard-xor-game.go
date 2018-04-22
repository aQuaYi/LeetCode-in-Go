package problem0810

func xorGame(nums []int) bool {
	n := len(nums)
	xor := 0
	for i := 0; i < n; i++ {
		xor ^= nums[i]
	}
	return xor == 0 || n%2 == 0
}

// Let’s discuss it if we add this condition.
// If xor == 0, Alice win directly.
// If xor != 0 and length of numbers is even, Alice will win.

// Beacause:
// All numbers won’t be the same. Otherwise xor will be equal to 0
// If all numbers are not the same, It means there are at least 2 different numbers.
// Alice can always erase a number different from current xor.
// So Alice won’t never lose this turn at this situation.
