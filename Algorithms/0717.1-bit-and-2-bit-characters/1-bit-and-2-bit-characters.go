package Problem0717

func isOneBitCharacter(bits []int) bool {
	n := len(bits)
	if n == 1 ||
		(n >= 2 && bits[n-2] == 0) ||
		(n == 3 && bits[n-3] == 1) ||
		(n >= 4 && bits[n-4] == 0 && bits[n-3] == 1 && bits[n-2] == 1) {
		return true
	}
	return false
}

func deleteTwoBits(bits []int) []int {
	res := make([]int, 0, len(bits))
	isJump := false
	for i := range bits {
		if isJump {
			isJump = false
			continue
		}

		if bits[i] == 1 {
			isJump = true
			continue
		}

		res = append(res, bits[i])
	}

	return res
}
