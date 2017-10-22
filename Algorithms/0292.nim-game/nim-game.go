package Problem0292

func canWinNim(n int) bool {
	if n <= 3 {
		return true
	}

	return !canWinNim(n-1) || !canWinNim(n-2) || !canWinNim(n-3)
}
