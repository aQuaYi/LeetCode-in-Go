package problem0991

func brokenCalc(X int, Y int) int {
	if X > Y {
		return X - Y
	}

	if Y%2 == 0 {
		return brokenCalc(X, Y/2) + 1
	}

	if X < Y {
		return brokenCalc(X*2, Y) + 1
	}

	return 0
}
