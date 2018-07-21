package problem0518

func change(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}

	if len(coins) == 0 {
		return 0
	}

	res := 0
	coin := coins[0]
	coins = coins[1:]
	t := 0

	for t <= amount {
		res += change(amount-t, coins)
		t += coin
	}

	return res
}
