package Problem0166

import (
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	minus := ""
	if numerator*denominator < 0 {
		minus = "-"
	}

	n, d := abs(numerator), abs(denominator)

	di := 0
	base := 0
	if n < d {
		for {
			n, base = next(n, d)
			di = di*base + n/d
			n %= d

			if n == 0 {
				break
			}

			if ds, ok := checkRepeat(di); ok {
				return ds
			}
		}

		return "0." + strconv.Itoa(di)
	}

	ds := ""
	ds = fractionToDecimal(n%d, d)
	return minus + strconv.Itoa(n/d) + ds[1:]
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func next(n, d int) (int, int) {
	base := 1
	for n < d {
		n *= 10
		base *= 10
	}
	return n, base
}

func checkRepeat(num int) (string, bool) {
	s := strconv.Itoa(num)

	n := len(s)

	start := n % 2
	for i := start; i <= n-2; i += 2 {
		i1 := (n + i) / 2
		if s[i:i1] == s[i1:] {
			return "0." + s[:i] + "(" + s[i:i1] + ")", true
		}
	}

	return "", false
}
