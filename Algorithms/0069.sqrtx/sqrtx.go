package Problem0069

func mySqrt(x int) int {
	res := 1
	for !(res*res <= x && x < (res+1)*(res+1)) {
		res = (res + x/res) / 2
	}

	return res
}
